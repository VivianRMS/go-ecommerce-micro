package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/VivianRMS/go-ecommerce-micro/account"
	"github.com/VivianRMS/go-ecommerce-micro/catalog"
	pb "github.com/VivianRMS/go-ecommerce-micro/order/gen/go/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	pb.UnimplementedOrderServiceServer
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		service:       s,
		accountClient: accountClient,
		catalogClient: catalogClient,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	// check if account exists from account service
	_, err := s.accountClient.GetAccount(ctx, r.AccountId)
	if err != nil {
		log.Println("Error getting account:", err)
		return nil, errors.New("account not found")
	}

	// get ordered products from catalog service
	productIDs := []string{}
	for _, p := range r.Products {
		productIDs = append(productIDs, p.ProductId)
	}
	orderedProducts, err := s.catalogClient.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		log.Println("Error getting products:", err)
		return nil, errors.New("products not found")
	}

	// construct order products
	products := []OrderedProduct{}
	for _, p := range orderedProducts {
		product := OrderedProduct{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Quantity:    0,
		}
		for _, rp := range r.Products {
			if rp.ProductId == p.ID {
				product.Quantity = rp.Quantity
				break
			}
		}
		if product.Quantity != 0 {
			products = append(products, product)
		}
	}

	// call order service to post order
	order, err := s.service.PostOrder(ctx, r.AccountId, products)
	if err != nil {
		log.Println("Error posting order:", err)
		return nil, errors.New("error posting order")
	}

	// make response order
	orderProto := &pb.Order{
		Id:         order.ID,
		AccountId:  order.AccountID,
		TotalPrice: order.TotalPrice,
		Products:   []*pb.Order_OrderProduct{},
	}
	orderProto.CreatedAt, _ = order.CreatedAt.MarshalBinary()
	for _, p := range order.Products {
		orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Quantity:    p.Quantity,
			Price:       p.Price,
		})
	}
	return &pb.PostOrderResponse{
		Order: orderProto,
	}, nil
}
func (s *grpcServer) GetOrdersForAccount(ctx context.Context, r *pb.GetOrdersForAccountRequest) (*pb.GetOrdersForAccountResponse, error) {
	accountOrders, err := s.service.GetOrdersForAccount(ctx, r.AccountId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	productIDMap := map[string]bool{}
	for _, o := range accountOrders {
		for _, p := range o.Products {
			productIDMap[p.ID] = true
		}
	}
	productIDs := []string{}
	for id := range productIDMap {
		productIDs = append(productIDs, id)
	}
	products, err := s.catalogClient.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		log.Println("Error getting products:", err)
		return nil, err
	}

	orders := []*pb.Order{}
	for _, o := range accountOrders {
		op := &pb.Order{
			Id:         o.ID,
			AccountId:  o.AccountID,
			TotalPrice: o.TotalPrice,
			Products:   []*pb.Order_OrderProduct{},
		}
		op.CreatedAt, _ = o.CreatedAt.MarshalBinary()
		//log.Println("get", op.CreatedAt)
		for _, product := range o.Products {
			for _, p := range products {
				if product.ID == p.ID {
					product.Name = p.Name
					product.Description = p.Description
					product.Price = p.Price
					break
				}
			}

			op.Products = append(op.Products, &pb.Order_OrderProduct{
				Id:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Quantity:    product.Quantity,
				Price:       product.Price,
			})
		}
		orders = append(orders, op)
	}
	return &pb.GetOrdersForAccountResponse{
		Orders: orders,
	}, nil
}

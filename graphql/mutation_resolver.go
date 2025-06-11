package main

import (
	"context"
	"github.com/VivianRMS/go-ecommerce-micro/order"
	"log"
	"time"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostAccount(ctx, in.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Account{
		ID:   a.ID,
		Name: a.Name,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, in ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.catalogClient.PostProduct(ctx, in.Name, in.Description, in.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Product{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, in OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var products []order.OrderedProduct
	for _, p := range in.Products {
		products = append(products, order.OrderedProduct{
			ID:       p.ID,
			Quantity: uint32(p.Quantity),
		})
	}

	o, err := r.server.orderClient.PostOrder(ctx, in.AccountID, products)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Order{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		TotalPrice: o.TotalPrice,
	}, nil
}

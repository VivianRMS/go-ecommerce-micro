// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: catalog/v1/catalog.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PostProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostProductRequest) Reset() {
	*x = PostProductRequest{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostProductRequest) ProtoMessage() {}

func (x *PostProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostProductRequest.ProtoReflect.Descriptor instead.
func (*PostProductRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *PostProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PostProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostProductResponse) Reset() {
	*x = PostProductResponse{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostProductResponse) ProtoMessage() {}

func (x *PostProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostProductResponse.ProtoReflect.Descriptor instead.
func (*PostProductResponse) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *PostProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Skip          uint64                 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take          uint64                 `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	Ids           []string               `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	Query         string                 `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductsRequest) GetSkip() uint64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetProductsRequest) GetTake() uint64 {
	if x != nil {
		return x.Take
	}
	return 0
}

func (x *GetProductsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetProductsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type GetProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	mi := &file_catalog_v1_catalog_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_catalog_v1_catalog_proto protoreflect.FileDescriptor

const file_catalog_v1_catalog_proto_rawDesc = "" +
	"\n" +
	"\x18catalog/v1/catalog.proto\x12\n" +
	"catalog.v1\"e\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\"`\n" +
	"\x12PostProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\"D\n" +
	"\x13PostProductResponse\x12-\n" +
	"\aproduct\x18\x01 \x01(\v2\x13.catalog.v1.ProductR\aproduct\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"C\n" +
	"\x12GetProductResponse\x12-\n" +
	"\aproduct\x18\x01 \x01(\v2\x13.catalog.v1.ProductR\aproduct\"d\n" +
	"\x12GetProductsRequest\x12\x12\n" +
	"\x04skip\x18\x01 \x01(\x04R\x04skip\x12\x12\n" +
	"\x04take\x18\x02 \x01(\x04R\x04take\x12\x10\n" +
	"\x03ids\x18\x03 \x03(\tR\x03ids\x12\x14\n" +
	"\x05query\x18\x04 \x01(\tR\x05query\"F\n" +
	"\x13GetProductsResponse\x12/\n" +
	"\bproducts\x18\x01 \x03(\v2\x13.catalog.v1.ProductR\bproducts2\x83\x02\n" +
	"\x0eCatalogService\x12P\n" +
	"\vPostProduct\x12\x1e.catalog.v1.PostProductRequest\x1a\x1f.catalog.v1.PostProductResponse\"\x00\x12M\n" +
	"\n" +
	"GetProduct\x12\x1d.catalog.v1.GetProductRequest\x1a\x1e.catalog.v1.GetProductResponse\"\x00\x12P\n" +
	"\vGetProducts\x12\x1e.catalog.v1.GetProductsRequest\x1a\x1f.catalog.v1.GetProductsResponse\"\x00BFZDgithub.com/VivianRMS/go-ecommerce-micro/catalog/gen/go/catalog/v1;pbb\x06proto3"

var (
	file_catalog_v1_catalog_proto_rawDescOnce sync.Once
	file_catalog_v1_catalog_proto_rawDescData []byte
)

func file_catalog_v1_catalog_proto_rawDescGZIP() []byte {
	file_catalog_v1_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_catalog_v1_catalog_proto_rawDesc), len(file_catalog_v1_catalog_proto_rawDesc)))
	})
	return file_catalog_v1_catalog_proto_rawDescData
}

var file_catalog_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_catalog_v1_catalog_proto_goTypes = []any{
	(*Product)(nil),             // 0: catalog.v1.Product
	(*PostProductRequest)(nil),  // 1: catalog.v1.PostProductRequest
	(*PostProductResponse)(nil), // 2: catalog.v1.PostProductResponse
	(*GetProductRequest)(nil),   // 3: catalog.v1.GetProductRequest
	(*GetProductResponse)(nil),  // 4: catalog.v1.GetProductResponse
	(*GetProductsRequest)(nil),  // 5: catalog.v1.GetProductsRequest
	(*GetProductsResponse)(nil), // 6: catalog.v1.GetProductsResponse
}
var file_catalog_v1_catalog_proto_depIdxs = []int32{
	0, // 0: catalog.v1.PostProductResponse.product:type_name -> catalog.v1.Product
	0, // 1: catalog.v1.GetProductResponse.product:type_name -> catalog.v1.Product
	0, // 2: catalog.v1.GetProductsResponse.products:type_name -> catalog.v1.Product
	1, // 3: catalog.v1.CatalogService.PostProduct:input_type -> catalog.v1.PostProductRequest
	3, // 4: catalog.v1.CatalogService.GetProduct:input_type -> catalog.v1.GetProductRequest
	5, // 5: catalog.v1.CatalogService.GetProducts:input_type -> catalog.v1.GetProductsRequest
	2, // 6: catalog.v1.CatalogService.PostProduct:output_type -> catalog.v1.PostProductResponse
	4, // 7: catalog.v1.CatalogService.GetProduct:output_type -> catalog.v1.GetProductResponse
	6, // 8: catalog.v1.CatalogService.GetProducts:output_type -> catalog.v1.GetProductsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_catalog_v1_catalog_proto_init() }
func file_catalog_v1_catalog_proto_init() {
	if File_catalog_v1_catalog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_catalog_v1_catalog_proto_rawDesc), len(file_catalog_v1_catalog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_v1_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_v1_catalog_proto_depIdxs,
		MessageInfos:      file_catalog_v1_catalog_proto_msgTypes,
	}.Build()
	File_catalog_v1_catalog_proto = out.File
	file_catalog_v1_catalog_proto_goTypes = nil
	file_catalog_v1_catalog_proto_depIdxs = nil
}

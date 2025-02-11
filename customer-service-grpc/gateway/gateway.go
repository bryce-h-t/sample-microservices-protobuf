package gateway

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/yarpc"
	pb "pl.piomin.services.grpc.customer.model"
)

type gateway struct {
	customerServiceClient pb.CustomersServiceClient
	tenancy string
}

// Interface defines methods for interacting with the CustomersService
type Interface interface {
	FindByPesel(ctx context.Context, pesel *wrappers.StringValue, opts ...yarpc.CallOption) (*pb.Customer, error)
	FindById(ctx context.Context, id *wrappers.Int32Value, opts ...yarpc.CallOption) (*pb.Customer, error)
	FindAll(ctx context.Context, empty *empty.Empty, opts ...yarpc.CallOption) (*pb.Customers, error)
	AddCustomer(ctx context.Context, customer *pb.Customer, opts ...yarpc.CallOption) (*pb.Customer, error)
}

// New creates a new instance of the CustomersService gateway
func New(client pb.CustomersServiceClient, tenancy string) Interface {
	return &gateway{
		customerServiceClient: client,
		tenancy: tenancy,
	}
}

// FindByPesel implements the FindByPesel RPC
func (g *gateway) FindByPesel(ctx context.Context, pesel *wrappers.StringValue, opts ...yarpc.CallOption) (*pb.Customer, error) {
	requestHeaders := []yarpc.CallOption{}
	requestHeaders = append(requestHeaders, opts...)
	return g.customerServiceClient.FindByPesel(ctx, pesel, requestHeaders...)
}

// FindById implements the FindById RPC
func (g *gateway) FindById(ctx context.Context, id *wrappers.Int32Value, opts ...yarpc.CallOption) (*pb.Customer, error) {
	requestHeaders := []yarpc.CallOption{}
	requestHeaders = append(requestHeaders, opts...)
	return g.customerServiceClient.FindById(ctx, id, requestHeaders...)
}

// FindAll implements the FindAll RPC
func (g *gateway) FindAll(ctx context.Context, empty *empty.Empty, opts ...yarpc.CallOption) (*pb.Customers, error) {
	requestHeaders := []yarpc.CallOption{}
	requestHeaders = append(requestHeaders, opts...)
	return g.customerServiceClient.FindAll(ctx, empty, requestHeaders...)
}

// AddCustomer implements the AddCustomer RPC
func (g *gateway) AddCustomer(ctx context.Context, customer *pb.Customer, opts ...yarpc.CallOption) (*pb.Customer, error) {
	requestHeaders := []yarpc.CallOption{}
	requestHeaders = append(requestHeaders, opts...)
	return g.customerServiceClient.AddCustomer(ctx, customer, requestHeaders...)
}

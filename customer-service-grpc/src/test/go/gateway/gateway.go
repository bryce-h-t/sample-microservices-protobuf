package gateway

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	customerPb "pl.piomin.services.grpc.customer.model"
	"github.com/stretchr/testify/require"
)

type gateway struct {
	customersClient customerPb.CustomersServiceClient
}

// Interface for talking to CustomersService endpoints
type Interface interface {
	FindByPesel(context.Context, *wrappers.StringValue) (*customerPb.Customer, error)
	FindById(context.Context, *wrappers.Int32Value) (*customerPb.Customer, error)
	FindAll(context.Context, *empty.Empty) (*customerPb.Customers, error)
	AddCustomer(context.Context, *customerPb.Customer) (*customerPb.Customer, error)
	GetCustomerID() int32
	GetCustomerPesel() string
}

// New creates a new gateway instance
func New(client customerPb.CustomersServiceClient) Interface {
	return &gateway{
		customersClient: client,
	}
}

// FindByPesel is RPC call to find customer by PESEL
func (g *gateway) FindByPesel(ctx context.Context, request *wrappers.StringValue) (*customerPb.Customer, error) {
	return g.customersClient.FindByPesel(ctx, request)
}

// FindById is RPC call to find customer by ID
func (g *gateway) FindById(ctx context.Context, request *wrappers.Int32Value) (*customerPb.Customer, error) {
	return g.customersClient.FindById(ctx, request)
}

// FindAll is RPC call to get all customers
func (g *gateway) FindAll(ctx context.Context, request *empty.Empty) (*customerPb.Customers, error) {
	return g.customersClient.FindAll(ctx, request)
}

// AddCustomer is RPC call to add a new customer
func (g *gateway) AddCustomer(ctx context.Context, request *customerPb.Customer) (*customerPb.Customer, error) {
	return g.customersClient.AddCustomer(ctx, request)
}

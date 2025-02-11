package flows_test

import (
    "context"
    "testing"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/wrappers"
    "go.uber.org/ctf"
    "customer-service-grpc/src/main/proto/flows"
    pb "pl.piomin.services.grpc.customer.model"
    "pl.piomin.services.grpc.customer.model/gateway"
)

// testGateway implements gateway.Interface for testing
type testGateway struct {
    customers map[int32]*pb.Customer
    nextID   int32
}

func (g *testGateway) AddCustomer(ctx context.Context, customer *pb.Customer) (*pb.Customer, error) {
    g.nextID++
    customer.Id = g.nextID
    g.customers[customer.Id] = customer
    return customer, nil
}

func (g *testGateway) FindById(ctx context.Context, id *wrappers.Int32Value) (*pb.Customer, error) {
    if customer, ok := g.customers[id.Value]; ok {
        return customer, nil
    }
    return nil, nil
}

func (g *testGateway) FindAll(ctx context.Context, _ *empty.Empty) (*pb.Customers, error) {
    customerList := make([]*pb.Customer, 0, len(g.customers))
    for _, customer := range g.customers {
        customerList = append(customerList, customer)
    }
    return &pb.Customers{Customers: customerList}, nil
}

func (g *testGateway) FindByPesel(ctx context.Context, pesel *wrappers.StringValue) (*pb.Customer, error) {
    return nil, nil // Not used in this test
}

func setupTestGateway() gateway.Interface {
    return &testGateway{
        customers: make(map[int32]*pb.Customer),
        nextID:    0,
    }
}

func TestCustomerValidationFlow(t *testing.T) {
    f := flows.New(&flows.Params{
        Entities: struct {
            CustomerGateway gateway.Interface
        }{
            CustomerGateway: setupTestGateway(),
        },
    })
    
    ctf.RunFlow(t, f)
}

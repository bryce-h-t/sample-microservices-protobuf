package flows

import (
    "go.uber.org/ctf"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/stretchr/testify/require"
    "pl.piomin.services.grpc.customer.model/gateway"
    pb "pl.piomin.services.grpc.customer.model"
    "customer-service-grpc/src/main/proto/actions/addcustomer"
    "customer-service-grpc/src/main/proto/actions/findbyid"
    "customer-service-grpc/src/main/proto/actions/findall"
)

const FlowName = "test://customer-service-grpc/flows/customer_validation"

type Params struct {
    Entities struct {
        CustomerGateway gateway.Interface
    }
}

func New(p *Params) *ctf.Flow {
    var customerID int32
    
    return ctf.NewFlowV2(FlowName,
        addcustomer.New(&addcustomer.Params{
            Entities: addcustomer.Entities{
                CustomerGateway: p.Entities.CustomerGateway,
            },
            RequestModifiers: []addcustomer.RequestModifier{
                func(t *ctf.T, e addcustomer.Entities, req *pb.Customer) {
                    req.FirstName = "John"
                    req.LastName = "Doe"
                    req.Email = "john.doe@example.com"
                },
            },
            Assertions: struct {
                Response []addcustomer.ResponseAssertion
            }{
                Response: []addcustomer.ResponseAssertion{
                    func(t *ctf.T, e addcustomer.Entities, res *pb.Customer, err error) {
                        customerID = res.Id
                        require.Greater(t, customerID, int32(0), "Customer ID should be positive")
                    },
                },
            },
        }),
        findbyid.New(&findbyid.Params{
            Entities: findbyid.Entities{
                CustomerGateway: p.Entities.CustomerGateway,
            },
            RequestModifiers: []findbyid.RequestModifier{
                func(t *ctf.T, e findbyid.Entities, req *wrappers.Int32Value) {
                    req.Value = customerID
                },
            },
            Assertions: struct {
                Response []findbyid.ResponseAssertion
            }{
                Response: []findbyid.ResponseAssertion{
                    func(t *ctf.T, e findbyid.Entities, res *pb.Customer, err error) {
                        require.Equal(t, "John", res.FirstName)
                        require.Equal(t, "Doe", res.LastName)
                    },
                },
            },
        }),
        findall.New(&findall.Params{
            Entities: findall.Entities{
                CustomerGateway: p.Entities.CustomerGateway,
            },
            Assertions: struct {
                Response []findall.ResponseAssertion
            }{
                Response: []findall.ResponseAssertion{
                    func(t *ctf.T, e findall.Entities, res *pb.Customers, err error) {
                        require.NotEmpty(t, res.Customers)
                        found := false
                        for _, c := range res.Customers {
                            if c.Id == customerID {
                                found = true
                                break
                            }
                        }
                        require.True(t, found, "Added customer should be in the list")
                    },
                },
            },
        }),
    )
}

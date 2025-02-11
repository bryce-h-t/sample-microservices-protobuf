package flows

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/require"
	"go.uber.org/ctf"
	"pl.piomin.services.grpc.customer.model/actions/addcustomer"
	"pl.piomin.services.grpc.customer.model/actions/findbyid"
	"pl.piomin.services.grpc.customer.model/actions/findall"
	gateway "pl.piomin.services.grpc.customer.model/gateway"
	pb "pl.piomin.services.grpc.customer.model"
)

const FlowName = "test://customer-service-grpc/flows/customer_lifecycle"

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
					req.Pesel = "12345678901"
				},
			},
			Assertions: struct{ Response []addcustomer.ResponseAssertion }{
				Response: []addcustomer.ResponseAssertion{
					func(t *ctf.T, e addcustomer.Entities, res *pb.Customer, err error) {
						require.NotZero(t, res.Id)
						customerID = res.Id
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
			Assertions: struct{ Response []findbyid.ResponseAssertion }{
				Response: []findbyid.ResponseAssertion{
					func(t *ctf.T, e findbyid.Entities, res *pb.Customer, err error) {
						require.Equal(t, customerID, res.Id)
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
			Assertions: struct{ Response []findall.ResponseAssertion }{
				Response: []findall.ResponseAssertion{
					func(t *ctf.T, e findall.Entities, res *pb.Customers, err error) {
						found := false
						for _, c := range res.Customers {
							if c.Id == customerID {
								found = true
								break
							}
						}
						require.True(t, found, "Added customer should exist in full customer list")
					},
				},
			},
		}),
	)
}

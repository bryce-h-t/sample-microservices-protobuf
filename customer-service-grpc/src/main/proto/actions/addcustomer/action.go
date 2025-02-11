package addcustomer

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/ctf"
	gateway "pl.piomin.services.grpc.customer.model/gateway"
	pb "pl.piomin.services.grpc.customer.model"
)

// actionName for this action
const (
	actionName = "test://customer-service-grpc/actions/addcustomer"
)

// RequestModifier is a modifier function which adds to the request before calling the action
type RequestModifier func(*ctf.T, Entities, *pb.Customer)

// ResponseAssertion is an assertion function to assert response
type ResponseAssertion func(*ctf.T, Entities, *pb.Customer, error)

// Entities are the entity inputs to the action
type Entities struct {
	CustomerGateway gateway.Interface
}

// Params is input required to make calls to customer-service::AddCustomer
type Params struct {
	Entities         Entities
	RequestModifiers []RequestModifier
	Assertions      struct {
		Response []ResponseAssertion
	}
}

// New triggers an action to call customer-service::AddCustomer
func New(p *Params) *ctf.Action {
	return ctf.NewActionV2(actionName,
		func(t *ctf.T) {
			request := &pb.Customer{}

			for _, modifier := range p.RequestModifiers {
				modifier(t, p.Entities, request)
			}

			res, err := p.Entities.CustomerGateway.AddCustomer(t.Context(), request)
			require.NoError(t, err, "No error should occur while calling customer-service::AddCustomer endpoint")
			require.NotNil(t, res, "response should not be nil while calling customer-service::AddCustomer endpoint")

			for _, assertion := range p.Assertions.Response {
				assertion(t, p.Entities, res, err)
			}
		},
	)
}

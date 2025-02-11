package addcustomer

import (
	"github.com/stretchr/testify/require"
	customerPb "pl.piomin.services.grpc.customer.model"
	"github.com/uber/ctf"
	"github.com/uber/ctf/customer-service-grpc/gateway"
)

const (
	actionName = "test://customer-service/actions/customer-service-grpc/addcustomer"
)

type RequestModifier func(*ctf.T, Entities, *customerPb.Customer)

type ResponseAssertion func(*ctf.T, Entities, *customerPb.Customer, error)

type Entities struct {
	CustomerGateway gateway.Interface
}

type Params struct {
	Entities         Entities
	RequestModifiers []RequestModifier
	Assertions      struct {
		Response []ResponseAssertion
	}
}

func New(p *Params) *ctf.Action {
	return ctf.NewActionV2(actionName,
		func(t *ctf.T) {
			request := &customerPb.Customer{}

			for _, modifier := range p.RequestModifiers {
				modifier(t, p.Entities, request)
			}

			res, err := p.Entities.CustomerGateway.AddCustomer(t.Context(), request)
			require.NoError(t, err, "No error should occur while calling customer-service::addcustomer endpoint")
			require.NotNil(t, res, "response should not be nil while calling customer-service::addcustomer endpoint")

			for _, assertion := range p.Assertions.Response {
				assertion(t, p.Entities, res, err)
			}
		},
	)
}

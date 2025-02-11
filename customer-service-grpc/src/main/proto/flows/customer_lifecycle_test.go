package flows

import (
	"testing"
	"go.uber.org/ctf"
	gateway "pl.piomin.services.grpc.customer.model/gateway"
)

func TestCustomerLifecycleFlow(t *testing.T) {
	customerGateway := &gateway.MockCustomerGateway{}
	
	flow := New(&Params{
		Entities: struct {
			CustomerGateway gateway.Interface
		}{
			CustomerGateway: customerGateway,
		},
	})

	ctf.RunFlow(t, flow)
}

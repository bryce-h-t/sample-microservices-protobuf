package flows

import (
	"go.uber.org/ctf"
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/addcustomer"
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/findbypesel"
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/findall"
	gateway "pl.piomin.services.grpc.customer.model/gateway"
)

const (
	FlowName         = "test://customer-service-grpc/flows/customer-lifecycle"
	testCustomerName = "John Doe"
	testCustomerPesel = "12345678901"
	testCustomerType = "individual"
)

type Params struct {
	Entities struct {
		CustomerGateway gateway.Interface
	}
}

func New(p *Params) *ctf.Flow {
	return ctf.NewFlowV2(FlowName,
		// Step 1: Add a new customer
		addcustomer.New(&addcustomer.Params{
			Entities: addcustomer.Entities{
				CustomerGateway: p.Entities.CustomerGateway,
			},
			RequestModifiers: []addcustomer.RequestModifier{
				addcustomer.WithName(testCustomerName),
				addcustomer.WithPesel(testCustomerPesel),
				addcustomer.WithType(testCustomerType),
			},
			Assertions: struct {
				Response []addcustomer.ResponseAssertion
			}{
				Response: []addcustomer.ResponseAssertion{
					addcustomer.AssertSuccessfulResponse(),
					addcustomer.AssertName(testCustomerName),
					addcustomer.AssertPesel(testCustomerPesel),
					addcustomer.AssertType(testCustomerType),
				},
			},
		}),

		// Step 2: Find customer by PESEL
		findbypesel.New(&findbypesel.Params{
			Entities: findbypesel.Entities{
				CustomerGateway: p.Entities.CustomerGateway,
			},
			RequestModifiers: []findbypesel.RequestModifier{
				findbypesel.WithPesel(testCustomerPesel),
			},
			Assertions: struct {
				Response []findbypesel.ResponseAssertion
			}{
				Response: []findbypesel.ResponseAssertion{
					findbypesel.AssertSuccessfulResponse(),
					findbypesel.AssertName(testCustomerName),
					findbypesel.AssertPesel(testCustomerPesel),
				},
			},
		}),

		// Step 3: Verify customer appears in full list
		findall.New(&findall.Params{
			Entities: findall.Entities{
				CustomerGateway: p.Entities.CustomerGateway,
			},
			Assertions: struct {
				Response []findall.ResponseAssertion
			}{
				Response: []findall.ResponseAssertion{
					findall.AssertSuccessfulResponse(),
					findall.AssertContainsCustomerWithName(testCustomerName),
					findall.AssertContainsCustomerWithPesel(testCustomerPesel),
				},
			},
		}),
	)
}

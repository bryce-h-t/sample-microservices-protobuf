package customermanagement

import (
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/addcustomer"
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/findbyid"
	"github.com/bryce-h-t/sample-microservices-protobuf/customer-service-grpc/src/main/proto/actions/findall"
	"github.com/bryce-h-t/ctf"
)

const FlowName = "CustomerManagementFlow"

type Params struct {
	Entities Entities
}

type Entities struct {
	CustomerGateway interface{}
	Customer        interface{}
}

func New(p *Params) ctf.Flow {
	return ctf.NewFlowV2(FlowName,
		addcustomer.New(&addcustomer.Params{
			CustomerGateway: p.Entities.CustomerGateway,
			Entities: addcustomer.Entities{
				Customer: p.Entities.Customer,
				RequestModifiers: []addcustomer.RequestModifier{
					addcustomer.BuildCustomerRequest(),
				},
				Assertions: addcustomer.Assertions{
					CustomerAssertions: []addcustomer.CustomerAssertions{
						addcustomer.ValidateCustomerCreated(),
					},
				},
			},
		}),
		findbyid.New(&findbyid.Params{
			CustomerGateway: p.Entities.CustomerGateway,
			Entities: findbyid.Entities{
				Customer: p.Entities.Customer,
				RequestModifiers: []findbyid.RequestModifier{
					findbyid.BuildFindByIdRequest(),
				},
				Assertions: findbyid.Assertions{
					CustomerAssertions: []findbyid.CustomerAssertions{
						findbyid.ValidateCustomerFound(),
					},
				},
			},
		}),
		findall.New(&findall.Params{
			CustomerGateway: p.Entities.CustomerGateway,
			Entities: findall.Entities{
				Customer: p.Entities.Customer,
				Assertions: findall.Assertions{
					CustomerAssertions: []findall.CustomerAssertions{
						findall.ValidateCustomersFound(),
					},
				},
			},
		}),
	)
}

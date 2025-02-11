package addcustomer

import (
	"github.com/uber/ctf"
	customerPb "pl.piomin.services.grpc.customer.model"
)

// nameRequestModifier sets the customer name in the request
func nameRequestModifier(name string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *customerPb.Customer) {
		request.Name = name
	}
}

// customerTypeRequestModifier sets the customer type in the request
func customerTypeRequestModifier(customerType string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *customerPb.Customer) {
		switch customerType {
		case "individual":
			request.Type = customerPb.Customer_INDIVIDUAL
		case "company":
			request.Type = customerPb.Customer_COMPANY
		default:
			request.Type = customerPb.Customer_INDIVIDUAL
		}
	}
}

// defaultRequestModifier sets default fields from entities
func defaultRequestModifier() RequestModifier {
	return func(t *ctf.T, entities Entities, request *customerPb.Customer) {
		request.Id = entities.CustomerGateway.GetCustomerID()
		request.Pesel = entities.CustomerGateway.GetCustomerPesel()
	}
}

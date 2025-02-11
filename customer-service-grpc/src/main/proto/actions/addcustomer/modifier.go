package addcustomer

import (
	"go.uber.org/ctf"
	pb "pl.piomin.services.grpc.customer.model"
)

// WithName sets the customer name in the request
func WithName(name string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *pb.Customer) {
		request.Name = name
	}
}

// WithPesel sets the customer PESEL in the request
func WithPesel(pesel string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *pb.Customer) {
		request.Pesel = pesel
	}
}

// WithType sets the customer type in the request
func WithType(customerType string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *pb.Customer) {
		switch customerType {
		case "individual":
			request.Type = pb.Customer_INDIVIDUAL
		case "company":
			request.Type = pb.Customer_COMPANY
		default:
			t.Fatalf("invalid customer type: %s", customerType)
		}
	}
}

// WithAccountNumber adds an account with the specified number to the customer's accounts list
func WithAccountNumber(number string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *pb.Customer) {
		account := &pb.Account{
			Number: number,
		}
		request.Accounts = append(request.Accounts, account)
	}
}

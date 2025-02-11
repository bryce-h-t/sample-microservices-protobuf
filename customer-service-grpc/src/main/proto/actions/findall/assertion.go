package findall

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/ctf"
	pb "pl.piomin.services.grpc.customer.model"
)

// AssertSuccessfulResponse asserts the basic response validity
func AssertSuccessfulResponse() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		require.NoError(t, err, "there should be no error while finding all customers")
		require.NotNil(t, response, "response should not be nil while finding all customers")
		require.NotNil(t, response.Customers, "customers list should not be nil")
	}
}

// AssertCustomerCount asserts the number of customers in the response
func AssertCustomerCount(expectedCount int) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		require.Len(t, response.Customers, expectedCount, "number of customers should match")
	}
}

// AssertContainsCustomerWithName asserts that a customer with the given name exists
func AssertContainsCustomerWithName(expectedName string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		found := false
		for _, customer := range response.Customers {
			if customer.Name == expectedName {
				found = true
				break
			}
		}
		require.True(t, found, "response should contain customer with name %s", expectedName)
	}
}

// AssertContainsCustomerWithPesel asserts that a customer with the given PESEL exists
func AssertContainsCustomerWithPesel(expectedPesel string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		found := false
		for _, customer := range response.Customers {
			if customer.Pesel == expectedPesel {
				found = true
				break
			}
		}
		require.True(t, found, "response should contain customer with PESEL %s", expectedPesel)
	}
}

// AssertContainsCustomerWithType asserts that a customer with the given type exists
func AssertContainsCustomerWithType(expectedType string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		var customerType pb.Customer_CustomerType
		switch expectedType {
		case "individual":
			customerType = pb.Customer_INDIVIDUAL
		case "company":
			customerType = pb.Customer_COMPANY
		default:
			t.Fatalf("invalid customer type: %s", expectedType)
		}
		
		found := false
		for _, customer := range response.Customers {
			if customer.Type == customerType {
				found = true
				break
			}
		}
		require.True(t, found, "response should contain customer with type %s", expectedType)
	}
}

// AssertContainsCustomerWithAccountNumber asserts that a customer with an account having the given number exists
func AssertContainsCustomerWithAccountNumber(expectedNumber string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customers, err error) {
		found := false
		for _, customer := range response.Customers {
			for _, account := range customer.Accounts {
				if account.Number == expectedNumber {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		require.True(t, found, "response should contain customer with account number %s", expectedNumber)
	}
}

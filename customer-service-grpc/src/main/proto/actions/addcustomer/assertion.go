package addcustomer

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/ctf"
	pb "pl.piomin.services.grpc.customer.model"
)

// AssertSuccessfulResponse asserts the basic response validity
func AssertSuccessfulResponse() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		require.NoError(t, err, "there should be no error while adding customer")
		require.NotNil(t, response, "response should not be nil while adding customer")
	}
}

// AssertName asserts the customer name matches expected value
func AssertName(expectedName string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		require.Equal(t, expectedName, response.Name, "customer name should match")
	}
}

// AssertPesel asserts the customer PESEL matches expected value
func AssertPesel(expectedPesel string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		require.Equal(t, expectedPesel, response.Pesel, "customer PESEL should match")
	}
}

// AssertType asserts the customer type matches expected value
func AssertType(expectedType string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		var customerType pb.Customer_CustomerType
		switch expectedType {
		case "individual":
			customerType = pb.Customer_INDIVIDUAL
		case "company":
			customerType = pb.Customer_COMPANY
		default:
			t.Fatalf("invalid customer type: %s", expectedType)
		}
		require.Equal(t, customerType, response.Type, "customer type should match")
	}
}

// AssertAccountNumber asserts that an account with the given number exists
func AssertAccountNumber(expectedNumber string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		found := false
		for _, account := range response.Accounts {
			if account.Number == expectedNumber {
				found = true
				break
			}
		}
		require.True(t, found, "customer should have account with number %s", expectedNumber)
	}
}

// AssertNoAccounts asserts that the customer has no accounts
func AssertNoAccounts() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *pb.Customer, err error) {
		require.Empty(t, response.Accounts, "customer should have no accounts")
	}
}

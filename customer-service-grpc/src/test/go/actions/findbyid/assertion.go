package findbyid

import (
	"github.com/stretchr/testify/require"
	"github.com/uber/ctf"
	customerPb "pl.piomin.services.grpc.customer.model"
)

// AssertCustomerName asserts the customer name matches expected value
func AssertCustomerName(expectedName string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while finding customer by ID")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, expectedName, response.Name, "customer name should match")
	}
}

// AssertCustomerPesel asserts the customer PESEL matches expected value
func AssertCustomerPesel(expectedPesel string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while finding customer by ID")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, expectedPesel, response.Pesel, "customer PESEL should match")
	}
}

// AssertCustomerType asserts the customer type matches expected value
func AssertCustomerType(expectedType string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while finding customer by ID")
		require.NotNil(t, response, "response should not be nil")
		var customerType customerPb.Customer_CustomerType
		switch expectedType {
		case "individual":
			customerType = customerPb.Customer_INDIVIDUAL
		case "company":
			customerType = customerPb.Customer_COMPANY
		default:
			customerType = customerPb.Customer_INDIVIDUAL
		}
		require.Equal(t, customerType, response.Type, "customer type should match")
	}
}

// AssertCustomerID asserts the customer ID matches expected value
func AssertCustomerID() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while finding customer by ID")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, entities.CustomerGateway.GetCustomerID(), response.Id, "customer ID should match")
	}
}

// AssertCustomerHasAccounts asserts the customer has the expected number of accounts
func AssertCustomerHasAccounts(expectedCount int) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while finding customer by ID")
		require.NotNil(t, response, "response should not be nil")
		require.Len(t, response.Accounts, expectedCount, "customer should have expected number of accounts")
	}
}

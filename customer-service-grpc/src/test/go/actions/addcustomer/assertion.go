package addcustomer

import (
	"github.com/stretchr/testify/require"
	"github.com/uber/ctf"
	customerPb "pl.piomin.services.grpc.customer.model"
)

// AssertCustomerName asserts the created customer name matches expected value
func AssertCustomerName(expectedName string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while adding customer")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, expectedName, response.Name, "created customer name should match")
	}
}

// AssertCustomerPesel asserts the created customer PESEL matches expected value
func AssertCustomerPesel() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while adding customer")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, entities.CustomerGateway.GetCustomerPesel(), response.Pesel, "created customer PESEL should match")
	}
}

// AssertCustomerType asserts the created customer type matches expected value
func AssertCustomerType(expectedType string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while adding customer")
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
		require.Equal(t, customerType, response.Type, "created customer type should match")
	}
}

// AssertCustomerID asserts the created customer ID matches expected value
func AssertCustomerID() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while adding customer")
		require.NotNil(t, response, "response should not be nil")
		require.Equal(t, entities.CustomerGateway.GetCustomerID(), response.Id, "created customer ID should match")
	}
}

// AssertCustomerHasNoAccounts asserts the created customer has no accounts initially
func AssertCustomerHasNoAccounts() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customer, err error) {
		require.NoError(t, err, "no error should occur while adding customer")
		require.NotNil(t, response, "response should not be nil")
		require.Empty(t, response.Accounts, "newly created customer should have no accounts")
	}
}

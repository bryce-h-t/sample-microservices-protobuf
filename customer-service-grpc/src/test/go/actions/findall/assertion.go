package findall

import (
	"github.com/stretchr/testify/require"
	"github.com/uber/ctf"
	customerPb "pl.piomin.services.grpc.customer.model"
)

// AssertCustomersNotEmpty asserts that the customers list is not empty
func AssertCustomersNotEmpty() ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customers, err error) {
		require.NoError(t, err, "no error should occur while finding all customers")
		require.NotNil(t, response, "response should not be nil")
		require.NotEmpty(t, response.Customers, "customers list should not be empty")
	}
}

// AssertCustomersCount asserts the total number of customers
func AssertCustomersCount(expectedCount int) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customers, err error) {
		require.NoError(t, err, "no error should occur while finding all customers")
		require.NotNil(t, response, "response should not be nil")
		require.Len(t, response.Customers, expectedCount, "customers count should match expected value")
	}
}

// AssertCustomerExistsByName asserts that a customer with the given name exists
func AssertCustomerExistsByName(expectedName string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customers, err error) {
		require.NoError(t, err, "no error should occur while finding all customers")
		require.NotNil(t, response, "response should not be nil")
		
		found := false
		for _, customer := range response.Customers {
			if customer.Name == expectedName {
				found = true
				break
			}
		}
		require.True(t, found, "customer with name %s should exist", expectedName)
	}
}

// AssertCustomerExistsByPesel asserts that a customer with the given PESEL exists
func AssertCustomerExistsByPesel(expectedPesel string) ResponseAssertion {
	return func(t *ctf.T, entities Entities, response *customerPb.Customers, err error) {
		require.NoError(t, err, "no error should occur while finding all customers")
		require.NotNil(t, response, "response should not be nil")
		
		found := false
		for _, customer := range response.Customers {
			if customer.Pesel == expectedPesel {
				found = true
				break
			}
		}
		require.True(t, found, "customer with PESEL %s should exist", expectedPesel)
	}
}

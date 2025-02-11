package findbypesel

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/uber/ctf"
)

// peselRequestModifier for setting PESEL in the request
func peselRequestModifier() RequestModifier {
	return func(t *ctf.T, entities Entities, request *wrappers.StringValue) {
		request.Value = entities.CustomerGateway.GetCustomerPesel()
	}
}

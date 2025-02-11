package findbyid

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/uber/ctf"
)

// customerIDRequestModifier for setting customer ID in the request
func customerIDRequestModifier() RequestModifier {
	return func(t *ctf.T, entities Entities, request *wrappers.Int32Value) {
		request.Value = entities.CustomerGateway.GetCustomerID()
	}
}

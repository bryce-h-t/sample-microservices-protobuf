package findbyid

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/ctf"
)

// criteriaRequestModifier for setting customer ID in the request
func criteriaRequestModifier() RequestModifier {
	return func(t *ctf.T, entities Entities, request *wrappers.Int32Value) {
		request.Value = int32(entities.CustomerGateway.Id)
	}
}

package findbypesel

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/ctf"
)

// WithPesel sets the PESEL value in the request
func WithPesel(pesel string) RequestModifier {
	return func(t *ctf.T, entities Entities, request *wrappers.StringValue) {
		request.Value = pesel
	}
}

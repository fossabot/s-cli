package policy

import (
	"skynx.io/s-cli/pkg/client/subnet"
)

func (api *API) Show() {
	s := subnet.GetSubnet(false)

	Output().Show(s)
}

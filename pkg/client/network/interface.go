package network

import "skynx.io/s-cli/pkg/client/network/output"

type Interface interface {
	List()
	Show()
	New()
	Update()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

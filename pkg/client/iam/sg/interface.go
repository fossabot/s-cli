package sg

import "skynx.io/s-cli/pkg/client/iam/sg/output"

type Interface interface {
	List()
	Show()
	Set()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

package project

import "skynx.io/s-cli/pkg/client/ops/project/output"

type Interface interface {
	List()
	Show()
	Create()
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

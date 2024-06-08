package role

import "skynx.io/s-cli/pkg/client/iam/role/output"

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

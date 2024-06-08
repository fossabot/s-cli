package user

import (
	"skynx.io/s-cli/pkg/client/iam/user/output"
)

type Interface interface {
	List()
	Show()
	Delete()
	SetPermissions()
	Enable()
	Disable()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

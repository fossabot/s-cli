package vs

import "skynx.io/s-cli/pkg/client/vs/output"

type Interface interface {
	List()
	Show()
	New()
	Update()
	Delete()
	AddVSAppSvc()
	DeleteVSAppSvc()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

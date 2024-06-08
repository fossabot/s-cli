package alert

import (
	"skynx.io/s-api-go/grpc/resources/events"
	"skynx.io/s-cli/pkg/client/alert/output"
)

type Interface interface {
	List()
	Show() *events.Alert
	Delete()
	NewNote(a *events.Alert)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

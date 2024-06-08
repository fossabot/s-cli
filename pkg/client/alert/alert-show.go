package alert

import "skynx.io/s-api-go/grpc/resources/events"

func (api *API) Show() *events.Alert {
	a := getAlert()

	Output().Show(a)

	return a
}

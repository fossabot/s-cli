package event

import "skynx.io/s-cli/pkg/client/event/output"

func Output() output.Interface {
	return &output.API{}
}

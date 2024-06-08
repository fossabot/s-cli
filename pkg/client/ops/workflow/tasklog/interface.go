package tasklog

import "skynx.io/s-cli/pkg/client/ops/workflow/tasklog/output"

type Interface interface {
	List()
	Show()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

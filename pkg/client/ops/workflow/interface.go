package workflow

import "skynx.io/s-cli/pkg/client/ops/workflow/output"

type Interface interface {
	List()
	Show()
	Create(yamlFile string)
	Update(yamlFile string)
	Delete()
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

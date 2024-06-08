package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/client/event"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(workflows map[string]*ops.Workflow) {
	output.SectionHeader("Ops: Workflows")
	output.TitleT1("Workflow List")

	t := table.New()
	t.Header(colors.Black("WORKFLOW NAME"), colors.Black(output.Fit("FAILURE RATE", 14)))

	for _, wf := range workflows {
		wfID := colors.DarkWhite(output.Fit(wf.Name, 48))
		if wf.EventMetrics != nil {
			t.AddRow(wfID, event.Output().FailureProbability(wf.EventMetrics))
		} else {
			t.AddRow(wfID, "n/a")
		}
	}

	t.Render()
	fmt.Println()
}

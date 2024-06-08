package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(taskLogs []*ops.TaskLog) {
	output.SectionHeader("Ops: TaskLogs")
	output.TitleT1("TaskLog List")

	t := table.New()
	t.Header(colors.Black("TASK"), colors.Black("TIMESTAMP"), colors.Black("TARGET NODE"))

	for _, tl := range taskLogs {
		tm := tl.Status.Timestamp
		t.AddRow(colors.DarkWhite(tl.TaskName), output.DatetimeMilli(tm), tl.NodeName)
	}

	t.Render()
	fmt.Println()
}

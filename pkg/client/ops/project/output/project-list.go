package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(projects map[string]*ops.Project) {
	output.SectionHeader("Ops: Projects")
	output.TitleT1("Project List")

	t := table.New()
	t.Header(colors.Black("PROJECT NAME"), colors.Black("DESCRIPTION"))

	for _, p := range projects {
		t.AddRow(colors.DarkWhite(output.Fit(p.Name, 24)), output.Fit(p.Description, 40))
	}

	t.Render()
	fmt.Println()
}

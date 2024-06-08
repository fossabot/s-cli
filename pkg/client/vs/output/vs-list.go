package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(vss map[string]*topology.VS) {
	output.SectionHeader("Virtual Servers")
	output.TitleT1("VS List")

	t := table.New()
	t.Header(colors.Black("VS ID"), colors.Black("NAME"), colors.Black("DESCRIPTION"))

	for _, vs := range vss {
		t.AddRow(vs.VSID, colors.DarkWhite(vs.Name), output.Fit(vs.Description, 32))
	}

	t.Render()
	fmt.Println()
}

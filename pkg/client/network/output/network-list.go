package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(networks map[string]*topology.Network) {
	output.SectionHeader("Networks")
	output.TitleT1("Network List")

	t := table.New()
	t.Header(colors.Black("NETWORK ID"), colors.Black("NETWORK CIDR"), colors.Black("DESCRIPTION"))

	for _, n := range networks {
		t.AddRow(n.NetID, colors.DarkWhite(n.NetworkCIDR), output.Fit(n.Description, 32))
	}

	t.Render()
	fmt.Println()
}

package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(subnets map[string]*topology.Subnet) {
	output.SectionHeader("Subnets")
	output.TitleT1("Subnet List")

	t := table.New()
	t.Header(colors.Black("SUBNET ID"), colors.Black("SUBNET CIDR"), colors.Black("DESCRIPTION"))

	for _, s := range subnets {
		if s.IPAM != nil {
			t.AddRow(s.SubnetID, colors.DarkWhite(s.IPAM.SubnetCIDR), output.Fit(s.Description, 32))
		}
	}

	t.Render()
	fmt.Println()
}

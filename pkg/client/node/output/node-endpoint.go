package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) ShowEndpoint(e *topology.Endpoint) {
	output.SectionHeader("Endpoint Details")
	output.TitleT1("Network Endpoint")

	t := table.New()

	t.AddRow(colors.Black("Endpoint"), colors.DarkWhite(e.EndpointID))
	t.AddRow(colors.Black("FQDN"), colors.DarkWhite(e.DNSName+".skynx.local"))
	t.AddRow(colors.Black("IPv4"), colors.DarkWhite(e.IPv4))
	t.AddRow(colors.Black("IPv6"), colors.DarkWhite(e.IPv6))

	t.Render()
	fmt.Println()
}

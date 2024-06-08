package output

import (
	"fmt"

	"skynx.io/s-api-go/grpc/resources/tenant"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/output/table"
	"skynx.io/s-lib/pkg/utils/colors"
)

func (api *API) List(tenants map[string]*tenant.Tenant) {
	output.SectionHeader("Tenants")
	output.TitleT1("Tenant List")

	t := table.New()
	t.Header(colors.Black("TENANT NAME"), colors.Black("DESCRIPTION"))

	for _, tenant := range tenants {
		t.AddRow(colors.DarkWhite(tenant.Name), output.Fit(tenant.Description, 48))
	}

	t.Render()
	fmt.Println()
}

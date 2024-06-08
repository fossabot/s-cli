package sg

import "skynx.io/s-cli/pkg/client/tenant"

func (api *API) Show() {
	tenantMap := tenant.Tenants()

	Output().Show(GetSecurityGroup(false), tenantMap)
}

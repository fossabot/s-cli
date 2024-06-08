package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/k8s/resource"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/k8s"
	"skynx.io/s-lib/pkg/k8s/config"
	"skynx.io/s-lib/pkg/utils/colors"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) ConnectService() {
	s := subnet.GetSubnet(false)
	if s == nil {
		msg.Alert("No subnet found.")
		msg.Alert("Please, configure at least one.")
		os.Exit(1)
	}

	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	mgw := api.validKubernetesGateway(s, api.getKubernetesGateways())

	ss := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesServicesAnnotations(), false)

	ss.Stop()

	if len(allIDs) == 0 {
		msg.Info("All services already connected")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := fmt.Sprintf("Connect to %s", s.SubnetID)
	if err := survey.AskOne(
		&survey.MultiSelect{
			Message:  selectMsg,
			Options:  allIDs,
			PageSize: 10,
		},
		&selectedIDs,
		survey.WithIcons(input.SurveySetIcons),
	); err != nil {
		status.Error(err, "Unable to get response")
	}

	ss = output.Spinner()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			ss.Stop()
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		annt := newAnnotations(r, s)
		if err := k8s.API(api.kubeConfig).Resources().Service().InsertAnnotations(r.Namespace, r.Name, annt); err != nil {
			ss.Stop()
			status.Error(err, "Unable to set kubernetes service annotation")
		}
	}

	ss.Stop()

	fmt.Println()

	api.Services()

	if !mgw {
		mgwNode := colors.DarkWhite(fmt.Sprintf("mgw-%s-%s-...", s.SubnetID, s.NetID))
		fmt.Printf(`A new skynx ingress gateway was created in the namespace %s.

If want to connect more kubernetes services to other subnets,
it is highly recommended that you use that namespace for your
skynx gateways to keep things tidy.

You will see the gateway as %s in
the node list.

It will route all the services you want to connect to %s
in your kubernetes cluster.

Enjoy :-)
`, colors.DarkWhite(config.NamespaceDefault), mgwNode, colors.DarkWhite(s.SubnetID))

		fmt.Println()
	}
}

func newAnnotations(r *resource.KubernetesResource, s *topology.Subnet) map[string]string {
	dnsName, ok := r.ServiceAnnotations["skynx.com/dnsName"]
	if !ok {
		dnsName = r.Name
	}
	ipv4, ok := r.ServiceAnnotations["skynx.com/ipv4"]
	if !ok {
		ipv4 = "auto"
	}

	return map[string]string{
		"skynx.com/account": s.AccountID,
		"skynx.com/tenant":  s.TenantID,
		"skynx.com/network": s.NetID,
		"skynx.com/subnet":  s.SubnetID,
		"skynx.com/dnsName": dnsName,
		"skynx.com/ipv4":    ipv4,
	}
}

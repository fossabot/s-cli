package k8s

import (
	"fmt"
	"os"
	"time"

	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-cli/pkg/client/k8s/resource"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/k8s"
	"skynx.io/s-lib/pkg/k8s/config"
	"skynx.io/s-lib/pkg/utils/colors"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) validKubernetesGateway(s *topology.Subnet, gateways map[string]*resource.KubernetesResource) bool {
	for _, sgw := range gateways {
		if sgw.Labels.AccountID == s.AccountID &&
			sgw.Labels.TenantID == s.TenantID &&
			sgw.Labels.NetID == s.NetID &&
			sgw.Labels.SubnetID == s.SubnetID {
			// kubernetes gateway found for this subnet
			return true
		}
	}

	msg.Infof("No skynx ingress gateway found to route the %s in your cluster.", colors.DarkWhite(s.SubnetID))

	fmt.Println()

	mgwMsg := "Want to create one?"
	if input.GetConfirm(mgwMsg, true) {
		ss := output.Spinner()
		api.createKubernetesGateway(s)
		ss.Stop()

		return false
	} else {
		fmt.Println()
		fmt.Println("Exiting..")
		fmt.Println()

		os.Exit(0)
	}

	return false
}

func (api *API) createKubernetesGateway(s *topology.Subnet) {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	port, err := k8s.API(api.kubeConfig).Resources().Service().GetNodePort()
	if err != nil {
		status.Error(err, "Unable to get node port")
	}

	r := &resource.KubernetesResource{
		Namespace: config.NamespaceDefault,
		Name:      fmt.Sprintf("mgw-%s-%d", s.SubnetID, time.Now().Unix()),
	}

	ni, err := r.GetGatewayNodeInstance(s, int32(port))
	if err != nil {
		status.Error(err, "Unable to get node instance")
	}

	if err := k8s.API(api.kubeConfig).Objects().Node().CreateGateway(ni); err != nil {
		status.Error(err, "Unable to create kubernetes resources")
	}
}

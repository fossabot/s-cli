package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-cli/pkg/client/k8s/resource"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/k8s"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) ConnectPod() {
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

	ss := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesPods(), false)

	ss.Stop()

	if len(allIDs) == 0 {
		msg.Info("All pods already connected")
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

		if r.Connected {
			continue
		}

		ni, err := r.GetPodNodeInstance(s)
		if err != nil {
			ss.Stop()
			status.Error(err, "Unable to get node instance")
		}

		switch r.KubernetesResourceType {
		case resource.KubernetesResourceTypeStatefulSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().ConnectStatefulSet(r.Namespace, r.Name, ni); err != nil {
				ss.Stop()
				status.Error(err, "Unable to connect kubernetes statefulSet")
			}
		case resource.KubernetesResourceTypeDeployment:
			if err := k8s.API(api.kubeConfig).Objects().Node().ConnectDeployment(r.Namespace, r.Name, ni); err != nil {
				ss.Stop()
				status.Error(err, "Unable to connect kubernetes deployment")
			}
		case resource.KubernetesResourceTypeDaemonSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().ConnectDaemonSet(r.Namespace, r.Name, ni); err != nil {
				ss.Stop()
				status.Error(err, "Unable to connect kubernetes daemonSet")
			}
		}
	}

	ss.Stop()

	fmt.Println()

	api.Pods()
}

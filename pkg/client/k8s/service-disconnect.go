package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-lib/pkg/k8s"
	"skynx.io/s-lib/pkg/utils/msg"
)

func (api *API) DisconnectService() {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	s := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesServicesAnnotations(), true)

	s.Stop()

	if len(allIDs) == 0 {
		msg.Info("No services connected found")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := "Disconnect from skynx"
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

	s = output.Spinner()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		annt := removedAnnotations()
		if err := k8s.API(api.kubeConfig).Resources().Service().RemoveAnnotations(r.Namespace, r.Name, annt); err != nil {
			status.Error(err, "Unable to remove kubernetes service annotations")
		}
	}

	s.Stop()

	fmt.Println()

	api.Services()
}

func removedAnnotations() map[string]string {
	return map[string]string{
		"skynx.com/account": "",
		"skynx.com/tenant":  "",
		"skynx.com/network": "",
		"skynx.com/subnet":  "",
		"skynx.com/dnsName": "",
		"skynx.com/ipv4":    "",
	}
}

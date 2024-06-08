package k8s

import (
	"skynx.io/s-cli/pkg/client/k8s/output"
)

type Interface interface {
	DeleteGateway()

	ConnectService()
	DisconnectService()

	ConnectPod()
	DisconnectPod()

	Services()
	Pods()
}
type API struct {
	kubeConfig []byte
}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}

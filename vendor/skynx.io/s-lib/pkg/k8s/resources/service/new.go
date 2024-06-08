package service

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-lib/pkg/k8s/config"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *corev1.Service {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ni.Node.KubernetesAttrs.Name,
			Namespace: ni.Node.KubernetesAttrs.Namespace,
			Labels:    config.NodeLabels(ni),
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"skynx-app":       appLabel.String(), // skynx-node
				"skynx-nodegroup": ni.Node.NodeGroupID,
			},
			Ports: []corev1.ServicePort{
				{
					Name:     "skynx-p2p-tcp",
					Protocol: corev1.ProtocolTCP,
					Port:     ni.Node.Agent.ExternalPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: ni.Node.Agent.Port,
					},
					NodePort: ni.Node.Agent.ExternalPort,
				},
				{
					Name:     "skynx-p2p-quic",
					Protocol: corev1.ProtocolUDP,
					Port:     ni.Node.Agent.ExternalPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: ni.Node.Agent.Port,
					},
					NodePort: ni.Node.Agent.ExternalPort,
				},
			},
			SessionAffinity:       corev1.ServiceAffinityClientIP,
			ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeLocal,
			// LoadBalancerIP:        ni.ExternalIPv4,
			// LoadBalancerSourceRanges: []string{
			// 	"0.0.0.0/0",
			// },
		},
	}
}

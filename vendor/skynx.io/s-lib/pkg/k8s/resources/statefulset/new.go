package statefulset

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"skynx.io/s-api-go/grpc/resources/topology"
	"skynx.io/s-lib/pkg/k8s/config"
	"skynx.io/s-lib/pkg/sx"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *appsv1.StatefulSet {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	configVolumeName := fmt.Sprintf("%s-config", appLabel.String())

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ni.Node.KubernetesAttrs.Name,
			Namespace: ni.Node.KubernetesAttrs.Namespace,
			Labels:    config.NodeLabels(ni),
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: ni.Node.KubernetesAttrs.Name,
			Replicas:    sx.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"skynx-app":       appLabel.String(), // skynx-node
					"skynx-nodegroup": ni.Node.NodeGroupID,
				},
			},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: config.NodeLabels(ni),
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: config.ServiceAccountView,
					Hostname:           ni.Node.NodeID,
					Containers: []corev1.Container{
						{
							Name:            ni.Node.KubernetesAttrs.Name,
							Image:           ni.Node.KubernetesAttrs.Image,
							ImagePullPolicy: corev1.PullAlways,
							SecurityContext: &corev1.SecurityContext{
								Privileged: sx.Bool(true), // only needed by sysctl to enable ipv6
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{
										"net_admin",
									},
								},
								RunAsUser:    sx.Int64(0),
								RunAsGroup:   sx.Int64(0),
								RunAsNonRoot: sx.Bool(false),
							},
							Args: []string{
								"start",
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "skynx-p2p-tcp",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: ni.Node.Agent.Port,
								},
								{
									Name:          "skynx-p2p-quic",
									Protocol:      corev1.ProtocolUDP,
									ContainerPort: ni.Node.Agent.Port,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "dev-net-tun",
									MountPath: "/dev/net/tun",
									ReadOnly:  true,
								},
								{
									Name:      configVolumeName,
									MountPath: "/etc/skynx",
									ReadOnly:  true,
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									// cpu: 30m
									corev1.ResourceCPU: *resource.NewMilliQuantity(30, resource.DecimalSI),
									// memory: 100Mi
									corev1.ResourceMemory: *resource.NewQuantity(100*1024*1024, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									// cpu: 15m
									corev1.ResourceCPU: *resource.NewMilliQuantity(15, resource.DecimalSI),
									// memory: 50Mi
									corev1.ResourceMemory: *resource.NewQuantity(50*1024*1024, resource.BinarySI),
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "dev-net-tun",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/dev/net/tun",
								},
							},
						},
						{
							Name: configVolumeName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: ni.Node.KubernetesAttrs.Name,
								},
							},
						},
					},
				},
			},
		},
	}
}

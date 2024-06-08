package resource

const (
	KubernetesResourceTypeService int = iota
	KubernetesResourceTypeStatefulSet
	KubernetesResourceTypeDeployment
	KubernetesResourceTypeDaemonSet
)

type KubernetesResource struct {
	KubernetesResourceType int
	Namespace              string
	Name                   string
	Connected              bool

	NetStatus NetStatus

	Labels             KubernetesLabels
	ServiceAnnotations map[string]string
}

type NetStatus struct {
	TenantID string
	NetID    string
	SubnetID string
}

type KubernetesLabels struct {
	NodeType  string
	AccountID string
	TenantID  string
	// NodeID      string
	NodeGroupID string
	NetID       string
	SubnetID    string
}

func (r *KubernetesResource) ParseLabelsPod(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "skynx-type":
			r.Labels.NodeType = v
		case "skynx-account":
			r.Labels.AccountID = v
		case "skynx-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		case "skynx-nodegroup":
			r.Labels.NodeGroupID = v
		case "skynx-network":
			// r.Labels.NetID = v
			r.NetStatus.NetID = v
		case "skynx-subnet":
			// r.Labels.SubnetID = v
			r.NetStatus.SubnetID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseLabelsGateway(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "skynx-type":
			r.Labels.NodeType = v
		case "skynx-account":
			r.Labels.AccountID = v
		case "skynx-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		// case "skynx-node":
		// 	r.Labels.NodeID = v
		case "skynx-nodegroup":
			r.Labels.NodeGroupID = v
		case "skynx-network":
			r.Labels.NetID = v
			r.NetStatus.NetID = v
		case "skynx-subnet":
			r.Labels.SubnetID = v
			r.NetStatus.SubnetID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 &&
		len(r.Labels.NetID) > 0 &&
		len(r.Labels.SubnetID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseServiceAnnotations(annotations map[string]string) {
	if annotations == nil {
		return
	}

	var hasTenant, hasNetwork, hasSubnet bool

	for k, v := range annotations {
		switch k {
		case "skynx.com/account":
			r.ServiceAnnotations[k] = v
		case "skynx.com/tenant":
			r.ServiceAnnotations[k] = v
			r.NetStatus.TenantID = v
			hasTenant = true
		case "skynx.com/network":
			r.ServiceAnnotations[k] = v
			r.NetStatus.NetID = v
			hasNetwork = true
		case "skynx.com/subnet":
			r.ServiceAnnotations[k] = v
			r.NetStatus.SubnetID = v
			hasSubnet = true
		case "skynx.com/dnsName":
			r.ServiceAnnotations[k] = v
		case "skynx.com/ipv4":
			r.ServiceAnnotations[k] = v
		}
	}

	if hasTenant && hasNetwork && hasSubnet {
		r.Connected = true
	}
}

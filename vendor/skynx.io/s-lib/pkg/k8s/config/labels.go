package config

import "skynx.io/s-api-go/grpc/resources/topology"

type AppLabel string

const (
	AppLabelController AppLabel = "skynx-controller"
	AppLabelNode       AppLabel = "skynx-node"
)

func (a AppLabel) String() string {
	return string(a)
}

func NodeLabels(ni *topology.NodeInstance) map[string]string {
	return map[string]string{
		"skynx-app":       AppLabelNode.String(),
		"skynx-type":      ni.Node.Type.String(),
		"skynx-account":   ni.Node.AccountID,
		"skynx-tenant":    ni.Node.TenantID,
		"skynx-nodegroup": ni.Node.NodeGroupID,
		"skynx-network":   ni.Node.Cfg.NetID,
		"skynx-subnet":    ni.Node.Cfg.SubnetID,
		// "skynx-node":      ni.Node.NodeID,
	}
}

package output

import "skynx.io/s-api-go/grpc/resources/topology"

type Interface interface {
	List(networks map[string]*topology.Network)
	Show(n *topology.Network)
}
type API struct{}

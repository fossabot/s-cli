package output

import "skynx.io/s-api-go/grpc/resources/topology"

type Interface interface {
	Show(s *topology.Subnet)
}
type API struct{}

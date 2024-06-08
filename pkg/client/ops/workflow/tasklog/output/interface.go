package output

import "skynx.io/s-api-go/grpc/resources/ops"

type Interface interface {
	List(operations []*ops.TaskLog)
	Show(tl *ops.TaskLog)
}
type API struct{}

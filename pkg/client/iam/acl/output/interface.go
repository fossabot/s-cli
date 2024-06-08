package output

import (
	"skynx.io/s-api-go/grpc/resources/iam"
)

type Interface interface {
	List(acls map[string]*iam.ACL)
	Show(acl *iam.ACL)
}
type API struct{}

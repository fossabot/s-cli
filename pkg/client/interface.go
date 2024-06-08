package client

import (
	"skynx.io/s-cli/pkg/client/account"
	"skynx.io/s-cli/pkg/client/alert"
	"skynx.io/s-cli/pkg/client/auth"
	"skynx.io/s-cli/pkg/client/vs"

	// "skynx.io/s-cli/pkg/client/command"
	"skynx.io/s-cli/pkg/client/iam/acl"
	// "skynx.io/s-cli/pkg/client/iam/role"
	// "skynx.io/s-cli/pkg/client/iam/sg"
	"skynx.io/s-cli/pkg/client/iam/user"
	"skynx.io/s-cli/pkg/client/k8s"
	"skynx.io/s-cli/pkg/client/network"
	"skynx.io/s-cli/pkg/client/node"
	"skynx.io/s-cli/pkg/client/ops/project"
	"skynx.io/s-cli/pkg/client/ops/workflow"
	"skynx.io/s-cli/pkg/client/ops/workflow/tasklog"
	"skynx.io/s-cli/pkg/client/policy"
	"skynx.io/s-cli/pkg/client/subnet"
	"skynx.io/s-cli/pkg/client/tenant"
)

func Auth() auth.Interface {
	return &auth.API{}
}

func Account() account.Interface {
	return &account.API{}
}

func Tenant() tenant.Interface {
	return &tenant.API{}
}

func Network() network.Interface {
	return &network.API{}
}

func Subnet() subnet.Interface {
	return &subnet.API{}
}

func Node() node.Interface {
	return &node.API{}
}

func VS() vs.Interface {
	return &vs.API{}
}

func NetworkPolicy() policy.Interface {
	return &policy.API{}
}

func ACL() acl.Interface {
	return &acl.API{}
}

/*
func Role() role.Interface {
	return &role.API{}
}

func SecurityGroup() sg.Interface {
	return &sg.API{}
}
*/

func User() user.Interface {
	return &user.API{}
}

func Project() project.Interface {
	return &project.API{}
}

func Workflow() workflow.Interface {
	return &workflow.API{}
}

func TaskLog() tasklog.Interface {
	return &tasklog.API{}
}

func Alert() alert.Interface {
	return &alert.API{}
}

func Kubernetes() k8s.Interface {
	return &k8s.API{}
}

// func Command() command.Interface {
// 	return &command.API{}
// }

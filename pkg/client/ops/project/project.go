package project

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/ops"
	"skynx.io/s-api-go/grpc/resources/resource"
	tenant_pb "skynx.io/s-api-go/grpc/resources/tenant"
	"skynx.io/s-cli/pkg/client/tenant"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-cli/pkg/vars"
	"skynx.io/s-lib/pkg/utils/msg"
)

var projectsMap map[string]*ops.Project = nil

func GetProject() *ops.Project {
	pl := projects()

	if len(pl) == 0 {
		msg.Info("No projects found")
		os.Exit(1)
	}

	var projectOptID string
	projectsOpts := make([]string, 0)
	projects := make(map[string]*ops.Project)

	for _, p := range pl {
		projectOptID = p.Name
		projectsOpts = append(projectsOpts, projectOptID)
		projects[projectOptID] = p
	}

	sort.Strings(projectsOpts)

	projectOptID = input.GetSelect("Project:", "", projectsOpts, survey.Required)

	vars.ProjectID = projects[projectOptID].ProjectID

	return projects[projectOptID]
}

func projects() map[string]*ops.Project {
	t := tenant.GetTenant()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	lr := &ops.ListProjectsRequest{
		Meta: &resource.ListRequest{},
		Tenant: &tenant_pb.TenantReq{
			AccountID: t.AccountID,
			TenantID:  t.TenantID,
		},
	}

	projects := make(map[string]*ops.Project)

	for {
		pl, err := nxc.ListProjects(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list projects")
		}

		for _, p := range pl.Projects {
			projects[p.ProjectID] = p
		}

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	return projects
}

/*
func validProjectName(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	projectName := val.(string)

	if projectsMap == nil {
		projectsMap = projects()
	}

	for _, p := range projectsMap {
		if p.Name == projectName {
			return fmt.Errorf("project %s already exist", projectName)
		}
	}

	return nil
}
*/

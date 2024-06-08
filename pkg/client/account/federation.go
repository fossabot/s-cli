package account

/*
import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/controller"
	location_pb "skynx.io/s-api-go/grpc/resources/location"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
	"skynx.io/s-cli/pkg/vars"
	"skynx.io/s-lib/pkg/utils/msg"
)

func selectFederation() *controller.FederationSelected {
	ll := selectLocation()

	var lOptions []string
	locations := make(map[string]*location_pb.Location)

	for _, l := range ll {
		lID := fmt.Sprintf("%s / %s", l.Region, l.Zone)
		lOptions = append(lOptions, lID)
		locations[lID] = l
	}

	sort.Strings(lOptions)

	if len(lOptions) == 0 {
		msg.Alert("No location found")
		os.Exit(1)
	}

	lID := vars.LocationID

	if len(lID) == 0 {
		lID = input.GetSelect("Location:", "", lOptions, survey.Required)
	}

	l, ok := locations[lID]
	if !ok {
		msg.Alert("Invalid selection")
		os.Exit(1)
	}

	vars.LocationID = lID

	nxc, grpcConn := grpc.GetManagerAPIClient(false)
	defer grpcConn.Close()

	s := output.Spinner()

	f, err := nxc.SelectFederation(context.TODO(), l)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get federation")
	}

	s.Stop()

	return f
}
*/

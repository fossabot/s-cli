package account

/*
import (
	"context"

	"skynx.io/s-api-go/grpc/resources/location"
	"skynx.io/s-api-go/grpc/resources/resource"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/status"
)

func selectLocation() []*location.Location {
	nxc, grpcConn := grpc.GetManagerAPIClient(false)
	defer grpcConn.Close()

	lr := &location.ListLocationsRequest{
		Meta: &resource.ListRequest{},
	}

	var locations []*location.Location
	var ll *location.Locations = nil
	var err error

	for ll == nil || len(lr.Meta.PageToken) > 0 {
		ll, err = nxc.SelectLocation(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list locations")
		}

		locations = append(locations, ll.Locations...)

		lr.Meta.PageToken = ll.Meta.NextPageToken
	}

	return locations
}
*/

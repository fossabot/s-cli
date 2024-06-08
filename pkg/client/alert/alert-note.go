package alert

import (
	"context"
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"skynx.io/s-api-go/grpc/resources/events"
	"skynx.io/s-cli/pkg/grpc"
	"skynx.io/s-cli/pkg/input"
	"skynx.io/s-cli/pkg/output"
	"skynx.io/s-cli/pkg/status"
)

func (api *API) NewNote(a *events.Alert) {
	if a == nil {
		a = getAlert()
	}

	nxc, grpcConn := grpc.GetMonitoringAPIClient()
	defer grpcConn.Close()

	acr := &events.AlertNewCommentRequest{
		AccountID: a.AccountID,
		AlertID:   a.AlertID,
		Comment: &events.AlertComment{
			Timestamp: time.Now().UnixMilli(),
			// UserID: userID,
			// UserEmail: userEmail,
			Text: input.GetMultiline("Note:", "", "", survey.Required),
		},
	}

	s := output.Spinner()

	_, err := nxc.NewAlertComment(context.TODO(), acr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to add alert note")
	}

	s.Stop()

	fmt.Printf(`

Your input has been sent!

Thanks for using skynx!

`)
}

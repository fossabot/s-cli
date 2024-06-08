package runtime

import (
	"time"

	"skynx.io/s-api-go/grpc/network/sxsp"
)

const (
	APIVersionV1 = "v1"
	Kind         = "Job"

	JobTypeCommand = "cmdJob"
	JobTypeNetwork = "networkJob"
	JobTypeIPAM    = "ipamJob"
)

const (
	JobPriorityLow    = "LOW"
	JobPriorityMedium = "MEDIUM"
	JobPriorityHigh   = "HIGH"
)

type JobSpec struct {
	ID        string      `json:"eventId"`
	NodeID    string      `json:"nodeId"`
	JobType   string      `json:"jobType"`
	Priority  string      `json:"priority"`
	Timestamp time.Time   `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

type Job struct {
	APIVersion string  `json:"apiVersion"`
	Kind       string  `json:"kind"`
	Spec       JobSpec `json:"spec"`
}
type Jobs []Job

func NewJob(id, nodeID, jobType, priority string) *Job {
	j := new(Job)
	j.APIVersion = APIVersionV1
	j.Kind = Kind
	j.Spec.ID = id
	j.Spec.NodeID = nodeID
	j.Spec.JobType = jobType
	j.Spec.Priority = priority
	switch jobType {
	case JobTypeCommand:
		j.Spec.Payload = new(sxsp.Payload)
	case JobTypeNetwork:
	}

	return j
}

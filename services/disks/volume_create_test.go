package disks

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeCreateRequest_BuildRequest(t *testing.T) {
	b := VolumeCreateBody{Name: "m"}
	dID := "volume_id"
	req := &VolumeCreateRequest{Body: b, ProjectID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(volumeCreateEndpoint, mocks.MockUrl, dID), b, t)
}

func TestVolumeCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"disk_id"}}`, http.StatusOK
				},
				Req:      &VolumeCreateRequest{ProjectID: "project_id"},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "disk_id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeCreateRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestVolumeCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeCreateRequest{}, t)
}

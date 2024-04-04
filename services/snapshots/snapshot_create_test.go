package snapshots

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestSnapshotCreateRequest_BuildRequest(t *testing.T) {
	b := SnapshotCreateBody{Name: "m"}
	ID := "id"
	req := &SnapshotCreateRequest{Body: b, ServerID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(snapCreateEndpoint, mocks.MockUrl, ID), b, t)
}

func TestSnapshotCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result": {"id": "id"}}`, http.StatusOK
				},
				Req:      &SnapshotCreateRequest{ServerID: "id"},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &SnapshotCreateRequest{ServerID: "id"},
			},
		},
	)
}

func TestSnapshotCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&SnapshotCreateRequest{}, t)
}

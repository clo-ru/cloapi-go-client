package snapshots

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestSnapshotRestoreRequest_BuildRequest(t *testing.T) {
	ID := "id"
	b := SnapshotRestoreBody{Name: "name"}
	req := &SnapshotRestoreRequest{SnapshotID: ID, Body: b}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(snapRestoreEndpoint, mocks.MockUrl, ID), b, t)
}

func TestSnapshotRestoreRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{

			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return `{"result":{"id": "id"}}`, http.StatusOK },
				Req:            &SnapshotRestoreRequest{SnapshotID: "id"},
				Expected:       &clo.ResponseCreated{Result: clo.IdResult{ID: "id"}},
				Actual:         &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &SnapshotRestoreRequest{SnapshotID: "id"},
			},
		},
	)
}

func TestSnapshotRestoreRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&SnapshotRestoreRequest{}, t)
}

package snapshots

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestSnapshotDetailRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &SnapshotDetailRequest{SnapshotID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(snapDetailEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestSnapshotDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"name":"snap","child_servers":["server_id"]}}`, http.StatusOK
				},
				Req:      &SnapshotDetailRequest{SnapshotID: "id"},
				Expected: &SnapshotDetailResponse{Result: Snapshot{Name: "snap", ChildServers: []string{"server_id"}}},
				Actual:   &SnapshotDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req:      &SnapshotDetailRequest{SnapshotID: "id"},
				Expected: nil,
			},
		},
	)
}

func TestSnapshotDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&SnapshotDetailRequest{SnapshotID: "id"}, t)
}

package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ServerListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(serverListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestServerListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerListRequest{}, t)
}

func TestServerList_Filtering(t *testing.T) {
	intTesting.FilterTest(&ServerListRequest{}, t)
}

func TestServerListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2, "result": [{"id": "first_item_id", "name": "host.com", "flavor":{"ram":3,"vcpus":2}},{"id": "second_item_id", "name": "host.com", "flavor":{"ram":2,"vcpus":3}}]}`, http.StatusOK
				},
				Req: &ServerListRequest{ProjectID: "project_id"},
				Expected: &ServerListResponse{
					Count: 2,
					Result: []Server{
						{
							ID:     "first_item_id",
							Name:   "host.com",
							Flavor: ServerFlavor{Ram: 3, Vcpus: 2},
						},
						{
							ID:     "second_item_id",
							Name:   "host.com",
							Flavor: ServerFlavor{Ram: 2, Vcpus: 3},
						},
					},
				},
				Actual: &ServerListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &ServerListRequest{ProjectID: "project_id"},
			},
		},
	)
}

package project

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestImageListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ImageListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(imageListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestImageListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2, "result": [{"id": "first_item_id", "name": "first_item_name", "operation_system":{"os_family":"debian"}},{"id": "second_item_id", "name": "second_item_name", "operation_system": null}]}`,
						http.StatusOK
				},
				Req: &ImageListRequest{ProjectID: "project_id"},
				Expected: &ImageListResponse{
					Count: 2,
					Result: []Image{
						{
							ID:              "first_item_id",
							Name:            "first_item_name",
							OperationSystem: &OperationSystem{OsFamily: "debian"},
						},
						{
							ID:              "second_item_id",
							Name:            "second_item_name",
							OperationSystem: nil,
						},
					},
				},
				Actual: &ImageListResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req:      &ImageListRequest{ProjectID: "project_id"},
				Expected: nil,
			},
		},
	)

}

func TestImageListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ImageListRequest{}, t)
}

func TestImageListRequest_Filtering(t *testing.T) {
	intTesting.FilterTest(&ImageListRequest{}, t)
}

func TestImageListPaginator_NextPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli

	mocks.BodyStringFunc = func() (string, int) {
		return "1", http.StatusOK
	}
	var cases = []struct {
		ShouldFail bool
		Name       string
		Expected   string
		Limit      int
		Offset     int
	}{
		{
			Name:     "Success",
			Limit:    3,
			Offset:   3,
			Expected: "limit=3&offset=3",
		},
		{
			Name:       "WrongLimit",
			ShouldFail: true,
			Limit:      2,
			Offset:     3,
			Expected:   "limit=3&offset=3",
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := &ImageListRequest{}
			res := &ImageListResponse{}
			pg := clo.NewPaginator(cli, req, c.Limit, c.Offset)

			err = pg.NextPage(context.Background(), res)
			if !c.ShouldFail {
				assert.Equal(t, c.Expected, httpCli.URL.RawQuery)
			} else {
				assert.NotEqual(t, c.Expected, httpCli.URL.RawQuery)
			}
		})
	}
}

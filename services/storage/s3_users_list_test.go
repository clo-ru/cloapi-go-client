package storage

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

func TestS3UserListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3UserListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(s3UserListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3UserListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count":2,"result":[{"name":"first_item_name","status":"ACTIVE"},{"name":"second_item_name","quotas":[{"max_size":3,"max_objects":2}]}]}`, http.StatusOK
				},
				Req: &S3UserListRequest{},
				Expected: &S3UserListResponse{
					Count: 2,
					Result: []S3User{
						{Name: "first_item_name", Status: "ACTIVE"},
						{Name: "second_item_name", Quotas: []QuotaInfo{{MaxSize: 3, MaxObjects: 2}}},
					},
				},
				Actual: &S3UserListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserListRequest{},
			},
		},
	)
}

func TestS3UserListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserListRequest{}, t)
}

func TestS3UserList_Filtering(t *testing.T) {
	intTesting.FilterTest(&S3UserListRequest{}, t)
}

func TestS3UserListPaginator_NextPage(t *testing.T) {
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
			req := &S3UserListRequest{}
			res := &S3UserListResponse{}
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

func TestS3UserListPaginator_lastPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli
	mocks.BodyStringFunc = func() (string, int) {
		return `{"count": 2, "result": [{"id": "first_item_id", "ptr": "host.com", "attached_to_server":{"id":"server_id"}},{"id": "second_item_id", "ptr": "host.com", "attached_to_server":{"id":"server_id"}}]}`,
			http.StatusOK
	}
	req := &S3UserListRequest{}
	res := &S3UserListResponse{}
	pg := clo.NewPaginator(cli, req, 3, 3)

	assert.Equal(t, false, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Nil(t, err)
	assert.Equal(t, true, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Equal(t, "no more pages", err.Error())
}

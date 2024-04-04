package sshkeys

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

func TestKeyPairListRequest_BuildRequest(t *testing.T) {
	projectId := "some_id"
	req := &KeyPairListRequest{ProjectID: projectId}

	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(keypairListEndpoint, mocks.MockUrl, projectId), nil, t)
}

func TestKeyPairListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2,"result": [{"name":"first_item_name","public_key":"first_pubkey"},{"name":"second_item_name","id":"second_item_id"}]}`,
						http.StatusOK
				},
				Req: &KeyPairListRequest{},
				Expected: &KeyPairListResponse{
					Count: 2,
					Result: []KeyPair{
						{
							Name:      "first_item_name",
							PublicKey: "first_pubkey",
						},
						{
							Name: "second_item_name",
							ID:   "second_item_id",
						},
					},
				},
				Actual: &KeyPairListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &KeyPairListRequest{},
				Expected:       nil,
			},
		},
	)
}

func TestKeyPairListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&KeyPairListRequest{}, t)
}

func TestKeyPairList_Filtering(t *testing.T) {
	intTesting.FilterTest(&KeyPairListRequest{}, t)
}

func TestKeyPairListPaginator_NextPage(t *testing.T) {
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
			req := &KeyPairDetailRequest{}
			res := &KeyPairListResponse{}
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

func TestKeyPairListPaginator_lastPage(t *testing.T) {
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
	req := &KeyPairDetailRequest{}
	res := &KeyPairListResponse{}
	pg := clo.NewPaginator(cli, req, 3, 3)

	assert.Equal(t, false, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Nil(t, err)
	assert.Equal(t, true, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Equal(t, "no more pages", err.Error())
}

package project

import (
	"github.com/clo-ru/cloapi-go-client/clo"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"sync"
	"testing"
)

func TestImageListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := ImageListRequest{
		ProjectID: ID,
	}
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("Bearer %s", mocks.MockAuthKey))
	h.Add("Content-type", "application/json")
	h.Add("X-Add-Some", "SomeHeaderValue")
	rawReq, e := req.buildRequest(context.Background(), map[string]interface{}{
		"auth_key": mocks.MockAuthKey,
		"base_url": mocks.MockUrl,
	})
	rawReq.Header = h
	assert.Nil(t, e)
	expReq, _ := http.NewRequestWithContext(
		context.Background(), http.MethodGet, mocks.MockUrl+fmt.Sprintf(imageListEndpoint, ID), nil,
	)
	expReq.Header = h
	assert.Equal(t, expReq, rawReq)
}

func TestImageListRequest_Make(t *testing.T) {
	httpCli := mocks.MockClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	var cases = []struct {
		Name           string
		ShouldFail     bool
		Req            ImageListRequest
		BodyStringFunc func() (string, int)
		Expected       ImageListResponse
	}{
		{
			Name: "Success",
			BodyStringFunc: func() (string, int) {
				return `{"count": 2, "results": [{"id": "first_item_id", "name": "first_item_name", "operation_system":{"os_family":"debian"}},{"id": "second_item_id", "name": "second_item_name", "operation_system":{"os_family":"centos"}}]}`,
					http.StatusOK
			},
			Req: ImageListRequest{
				ProjectID: "project_id",
			},
			Expected: ImageListResponse{
				Count: 2,
				Results: []ImageListItem{
					{
						ID:   "first_item_id",
						Name: "first_item_name",
						OperationSystem: OperationSystemItem{
							OsFamily: "debian",
						},
					},
					{
						ID:   "second_item_id",
						Name: "second_item_name",
						OperationSystem: OperationSystemItem{
							OsFamily: "centos",
						},
					},
				},
			},
		},
		{
			Name:       "WrongCount",
			ShouldFail: true,
			BodyStringFunc: func() (string, int) {
				return `{"count": 2, "results": [{"id": "first_item_id", "name": "first_item_name", "operation_system":{"os_family":"debian"}},{"id": "second_item_id", "name": "second_item_name", "operation_system":{"os_family":"centos"}}]}`,
					http.StatusOK
			},
			Req: ImageListRequest{
				ProjectID: "project_id",
			},
			Expected: ImageListResponse{
				Count: 1,
				Results: []ImageListItem{
					{
						ID:   "first_item_id",
						Name: "first_item_name",
						OperationSystem: OperationSystemItem{
							OsFamily: "debian",
						},
					},
					{
						ID:   "second_item_id",
						Name: "second_item_name",
						OperationSystem: OperationSystemItem{
							OsFamily: "centos",
						},
					},
				},
			},
		},
		{
			Name:       "Error",
			ShouldFail: true,
			BodyStringFunc: func() (string, int) {
				return "", http.StatusInternalServerError
			},
			Req: ImageListRequest{
				ProjectID: "project_id",
			},
			Expected: ImageListResponse{Results: []ImageListItem{}},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			mocks.BodyStringFunc = c.BodyStringFunc
			res, e := c.Req.Make(context.Background(), &cli)
			if !c.ShouldFail {
				assert.Nil(t, e)
				assert.Equal(t, c.Expected, res)
			} else {
				assert.NotEqual(t, c.Expected, res)
			}
		})
	}
}

func TestImageListRequest_MakeRetry(t *testing.T) {
	retry := 5
	erCode := http.StatusInternalServerError
	httpCli := mocks.RequestDebugClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	mocks.BodyStringFunc = func() (string, int) {
		return "", erCode
	}
	grNum := 4
	wg := sync.WaitGroup{}
	for n := 0; n < grNum; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := ImageListRequest{}
			req.WithRetry(retry, 0)
			_, _ = req.Make(context.Background(), &cli)
		}()
	}
	wg.Wait()
	assert.Equal(t, retry*grNum, httpCli.RequestCount)
}

func TestImageListPaginator_NextPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	mocks.BodyStringFunc = func() (string, int) {
		return "1", http.StatusOK
	}
	var cases = []struct {
		ShouldFail bool
		Name       string
		Expected   string
		PGOptions  PaginatorOptions
	}{
		{
			Name: "Success",
			PGOptions: PaginatorOptions{
				Limit:  3,
				Offset: 3,
			},
			Expected: "limit=3&offset=3",
		},
		{
			Name:       "WrongLimit",
			ShouldFail: true,
			PGOptions: PaginatorOptions{
				Limit:  2,
				Offset: 3,
			},
			Expected: "limit=3&offset=3",
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := ImageListRequest{
				ProjectID: "id",
			}
			pg, e := NewImageListPaginator(&cli, req, c.PGOptions)
			assert.Nil(t, e)
			_, e = pg.NextPage(context.Background())
			if !c.ShouldFail {
				assert.Equal(t, c.Expected, httpCli.URL.RawQuery)
			} else {
				assert.NotEqual(t, c.Expected, httpCli.URL.RawQuery)
			}
		})
	}
}

func TestImageListPaginator_lastPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	req := ImageListRequest{
		ProjectID: "id",
	}
	mocks.BodyStringFunc = func() (string, int) {
		return `{"count": 2, "results": [{"id": "first_item_id", "ptr": "host.com", "attached_to_server":{"id":"server_id"}},{"id": "second_item_id", "ptr": "host.com", "attached_to_server":{"id":"server_id"}}]}`,
			http.StatusOK
	}
	pg, e := NewImageListPaginator(&cli, req, PaginatorOptions{
		Limit:  3,
		Offset: 3,
	})
	assert.Nil(t, e)
	assert.Equal(t, false, pg.lastPage)

	_, e = pg.NextPage(context.Background())
	assert.Nil(t, e)
	assert.Equal(t, true, pg.lastPage)

	_, e = pg.NextPage(context.Background())
	assert.Equal(t, "no more pages", e.Error())
}

func TestImageList_Filtering(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	mocks.BodyStringFunc = func() (string, int) {
		return "1", http.StatusOK
	}
	var cases = []struct {
		ShouldFail   bool
		Name         string
		OrderFields  []string
		FilterFields []FilteringField
		RawExpected  map[string][]string
	}{
		{
			Name: "Success",
			FilterFields: []FilteringField{
				{
					FieldName: "field_gt",
					Condition: "gt",
					Value:     "3",
				},
				{
					FieldName: "field_in",
					Condition: "in",
					Value:     "2,3,4",
				},
				{
					FieldName: "field_range",
					Condition: "range",
					Value:     "2:3",
				},
			},
			OrderFields: []string{
				"field3", "-field4",
			},
			RawExpected: map[string][]string{
				"field_gt__gt":       {"3"},
				"field_in__in":       {"2,3,4"},
				"field_range__range": {"2:3"},
				"order":              {"field3", "-field4"},
			},
		},
		{
			Name:       "WrongCondition",
			ShouldFail: true,
			FilterFields: []FilteringField{
				{
					FieldName: "field_gt",
					Condition: "gt",
					Value:     "3",
				},
				{
					FieldName: "field_in",
					Condition: "in",
					Value:     "2,3,4",
				},
				{
					FieldName: "field_range",
					Condition: "range",
					Value:     "2:3",
				},
			},
			OrderFields: []string{
				"field3", "-field4",
			},
			RawExpected: map[string][]string{
				"field_gt__gt":       {"2"},
				"field_in__in":       {"2,3,4"},
				"field_range__range": {"2:3"},
				"order":              {"field3", "-field4"},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var params url.Values
			params = c.RawExpected
			expected := params.Encode()
			req := ImageListRequest{
				ProjectID: "id",
			}
			for _, ff := range c.FilterFields {
				req.FilterBy(ff)
			}
			for _, of := range c.OrderFields {
				req.OrderBy(of)
			}
			_, _ = req.Make(context.Background(), &cli)
			if !c.ShouldFail {
				assert.Equal(t, expected, httpCli.URL.RawQuery)
			} else {
				assert.NotEqual(t, expected, httpCli.URL.RawQuery)
			}
		})
	}
}

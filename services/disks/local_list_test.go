package disks

import (
	"github.com/clo-ru/cloapi-go-client/clo"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestLocalListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := LocalListRequest{
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
		context.Background(), http.MethodGet, mocks.MockUrl+fmt.Sprintf(localListEndpoint, ID), nil,
	)
	expReq.Header = h
	assert.Equal(t, expReq, rawReq)
}

func TestLocalListRequest_Make(t *testing.T) {
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
		Req            LocalListRequest
		BodyStringFunc func() (string, int)
		Expected       LocalListResponse
	}{
		{
			Name: "Success",
			BodyStringFunc: func() (string, int) {
				return `{"count": 2, "results": [{"id": "first_item_id", "name": "first_item_name"},{"id": "second_item_id", "name": "second_item_name"}]}`, http.StatusOK
			},
			Req: LocalListRequest{
				ProjectID: "project_id",
			},
			Expected: LocalListResponse{
				Count: 2,
				Results: []ResponseItem{
					{
						ID:   "first_item_id",
						Name: "first_item_name",
					},
					{
						ID:   "second_item_id",
						Name: "second_item_name",
					},
				},
			},
		},
		{
			Name:       "WrongCount",
			ShouldFail: true,
			BodyStringFunc: func() (string, int) {
				return `{"count": 2, "results": [{"id": "first_item_id", "name": "first_item_name"},{"id": "second_item_id", "name": "second_item_name"}]}`,
					http.StatusOK
			},
			Req: LocalListRequest{
				ProjectID: "project_id",
			},
			Expected: LocalListResponse{
				Count: 1,
				Results: []ResponseItem{
					{
						ID:   "first_item_id",
						Name: "first_item_name",
					},
					{
						ID:   "second_item_id",
						Name: "second_item_name",
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
			Req: LocalListRequest{
				ProjectID: "project_id",
			},
			Expected: LocalListResponse{Results: []ResponseItem{}},
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

func TestLocalListRequest_MakeRetry(t *testing.T) {
	retry := 2
	erCode := http.StatusInternalServerError
	httpCli := mocks.MockClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	req := LocalListRequest{
		ProjectID: "id",
	}
	req.WithRetry(retry, 0)
	mocks.BodyStringFunc = func() (string, int) {
		return "", erCode
	}
	_, e := req.Make(context.Background(), &cli)
	assert.Equal(t, http.StatusText(erCode), e.Error())
	var er []error
	for i := 0; i < retry; i++ {
		er = append(er, errors.New(http.StatusText(erCode)))
	}
}

func TestLocalListPaginator_NextPage(t *testing.T) {
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
			req := LocalListRequest{
				ProjectID: "id",
			}
			pg, e := NewLocalListPaginator(&cli, &req, c.PGOptions)
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

func TestLocalListPaginator_lastPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cli := clo.ApiClient{
		HttpClient: &httpCli,
		Options: map[string]interface{}{
			"auth_key": "secret",
			"base_url": "https://clo.ru",
		},
	}
	req := LocalListRequest{
		ProjectID: "id",
	}
	mocks.BodyStringFunc = func() (string, int) {
		return `{"count": 2, "results": [{"id": "first_item_id", "name": "first_item_name"},{"id": "second_item_id", "name": "second_item_name"}]}`,
			http.StatusOK
	}
	pg, e := NewLocalListPaginator(&cli, &req, PaginatorOptions{
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

func TestLocalList_Filtering(t *testing.T) {
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
				"field_gt__gt":       {"3"},
				"field_in":           {"2,3,4"},
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
			req := LocalListRequest{
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

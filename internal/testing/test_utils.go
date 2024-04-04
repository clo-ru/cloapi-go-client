package testing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func FilterTest(request clo.FilterableRequest, t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli

	var cases = []struct {
		ShouldFail   bool
		Name         string
		OrderFields  []string
		FilterFields []clo.FilteringField
		RawExpected  map[string][]string
	}{
		{
			Name: "Success",
			FilterFields: []clo.FilteringField{
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
			FilterFields: []clo.FilteringField{
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
			for _, ff := range c.FilterFields {
				request.FilterBy(ff)
			}
			for _, of := range c.OrderFields {
				request.OrderBy(of)
			}
			cli.DoRequest(context.Background(), request, nil)
			if !c.ShouldFail {
				assert.Equal(t, expected, httpCli.URL.RawQuery)
			} else {
				assert.NotEqual(t, expected, httpCli.URL.RawQuery)
			}
		})
	}
}

func ConcurrentRetryTest(request clo.RequestInt, t *testing.T) {
	retry := 5
	grNum := 1000
	erCode := http.StatusInternalServerError

	httpCli := mocks.MockClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli
	mocks.BodyStringFunc = func() (string, int) { return "", erCode }
	request.WithRetry(retry, 0)
	wg := sync.WaitGroup{}
	for n := 0; n < grNum; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cli.DoRequest(context.Background(), request, nil)
		}()
	}
	wg.Wait()
	assert.Equal(t, retry*grNum, httpCli.RequestCount())
}

func BuildTest(request clo.RequestInt, method string, url string, expectedBody any, t *testing.T) {
	h := http.Header{}
	h.Add("X-Add-Some", "SomeHeaderValue")
	request.WithHeaders(h)
	rawReq, e := request.Build(context.Background(), mocks.MockUrl, mocks.MockAuthKey)
	assert.Nil(t, e)

	var expReq *http.Request
	if expectedBody != nil {
		body := new(bytes.Buffer)
		json.NewEncoder(body).Encode(expectedBody)
		expReq, _ = http.NewRequestWithContext(context.Background(), method, url, body)
	} else {
		expReq, _ = http.NewRequestWithContext(context.Background(), method, url, nil)
	}

	expectedH := http.Header{}
	expectedH.Add("Authorization", fmt.Sprintf("Bearer %s", mocks.MockAuthKey))
	expectedH.Add("Content-type", "application/json")
	expectedH.Add("X-Add-Some", "SomeHeaderValue")

	expReq.Header = expectedH
	assert.Equal(t, expReq.Method, rawReq.Method)
	assert.Equal(t, expReq.URL, rawReq.URL)
	assert.Equal(t, expReq.Body, rawReq.Body)
	assert.Equal(t, expReq.Header, rawReq.Header)
	assert.Equal(t, expReq.Header, rawReq.Header)
}

type DoTestCase struct {
	Name           string
	ShouldFail     bool
	CheckError     bool
	Req            clo.RequestInt
	BodyStringFunc func() (string, int)
	Expected       any
	Actual         any
}

func (c *DoTestCase) TestCase(t *testing.T, cli *clo.ApiClient) {
	if reflect.TypeOf(c.Expected) != reflect.TypeOf(c.Actual) {
		t.Errorf("Invalid case, Expected and Actual must be same type")
	}
	t.Run(c.Name, func(t *testing.T) {
		mocks.BodyStringFunc = c.BodyStringFunc
		err := cli.DoRequest(context.Background(), c.Req, c.Actual)
		if !c.ShouldFail {
			assert.Nil(t, err)
			assert.Equal(t, c.Expected, c.Actual)
		} else {
			if c.CheckError {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, c.Expected, c.Actual)
			}
		}
	})
}

func TestDoRequestCases(t *testing.T, cases []DoTestCase) {
	httpCli := mocks.MockClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli

	for _, c := range cases {
		c.TestCase(t, cli)
	}
}

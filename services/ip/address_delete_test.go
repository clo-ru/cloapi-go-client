package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestAddressDeleteRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := AddressDeleteRequest{
		AddressID: ID,
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
		context.Background(), http.MethodDelete, mocks.MockUrl+fmt.Sprintf(addressDeleteEndpoint, ID), nil,
	)
	expReq.Header = h
	assert.Equal(t, expReq, rawReq)
}

func TestAddressDeleteRequest_Make(t *testing.T) {
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
		CheckError     bool
		Req            AddressDeleteRequest
		BodyStringFunc func() (string, int)
	}{
		{
			Name: "Success",
			BodyStringFunc: func() (string, int) {
				return "1",
					http.StatusOK
			},
			Req: AddressDeleteRequest{
				AddressID: "id",
			},
		},
		{
			Name:       "Error",
			ShouldFail: true,
			CheckError: true,
			BodyStringFunc: func() (string, int) {
				return "", http.StatusInternalServerError
			},
			Req: AddressDeleteRequest{
				AddressID: "id",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			mocks.BodyStringFunc = c.BodyStringFunc
			e := c.Req.Make(context.Background(), &cli)
			if !c.ShouldFail {
				assert.Nil(t, e)
			} else {
				assert.NotNil(t, e)
			}
		})
	}
}

func TestAddressDeleteRequest_MakeRetry(t *testing.T) {
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
			req := AddressDeleteRequest{}
			req.WithRetry(retry, 0)
			_ = req.Make(context.Background(), &cli)
		}()
	}
	wg.Wait()
	assert.Equal(t, retry*grNum, httpCli.RequestCount)
}

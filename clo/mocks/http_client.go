package mocks

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
)

const (
	MockUrl     = "https://clo.ru"
	MockAuthKey = "supersecret"
)

var BodyStringFunc func() (string, int)

type MockClient struct {
	requestCount atomic.Int32
}

func (mc *MockClient) Do(req *http.Request) (*http.Response, error) {
	mc.requestCount.Add(1)
	var bodyString string
	var statusCode int
	var body io.ReadCloser = nil
	bodyString, statusCode = BodyStringFunc()
	//if len(bodyString) != 0 {
	//
	//}
	body = ioutil.NopCloser(bytes.NewReader([]byte(bodyString)))
	return &http.Response{StatusCode: statusCode, Body: body}, nil
}

func (mc *MockClient) RequestCount() int {
	return int(mc.requestCount.Load())
}

// RequestDebugClient is useful when you want to discover the passed URL and parameters
type RequestDebugClient struct {
	requestCount atomic.Int32
	URL          url.URL
	Headers      http.Header
	Body         []byte
}

func (rdc *RequestDebugClient) Do(req *http.Request) (*http.Response, error) {
	rdc.requestCount.Add(1)
	rdc.URL = *req.URL
	rdc.Headers = req.Header
	var e error
	if req.Body != nil {
		if rdc.Body, e = io.ReadAll(req.Body); e != nil {
			return nil, e
		}
		if e = req.Body.Close(); e != nil {
			return nil, e
		}
	}
	var bodyString string
	var statusCode int
	var body io.ReadCloser = nil
	bodyString, statusCode = BodyStringFunc()
	body = ioutil.NopCloser(bytes.NewReader([]byte(bodyString)))
	//if len(bodyString) != 0 {
	//	body = ioutil.NopCloser(bytes.NewReader([]byte(bodyString)))
	//}
	return &http.Response{StatusCode: statusCode, Body: body}, nil
}
func (rdc *RequestDebugClient) RequestCount() int {
	return int(rdc.requestCount.Load())
}
func CheckHeaders(headers map[string][]string) bool {
	if _, ok := headers["Authorization"]; !ok {
		return false
	}
	if _, ok := headers["Content-Type"]; !ok {
		return false
	}
	return true
}

package clo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type QueryParam map[string][]string

type RequestInt interface {
	RetryCount() int
	RetryDelay() time.Duration
	WithQueryParams(param QueryParam)
	WithRetry(retry int, timeout time.Duration)
	WithHeaders(headers http.Header)
	WithLog(l Logger)

	OnRetry(response *http.Response, err error, retryCount int)
	Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error)
}

type FilterableRequest interface {
	RequestInt
	OrderBy(of string)
	FilterBy(ff FilteringField)
}

type PaginatedRequest interface {
	RequestInt
	OrderBy(of string)
	FilterBy(ff FilteringField)
}

type Request struct {
	retry        int
	retryTimeout time.Duration
	log          Logger
	headers      http.Header
	queryParams  QueryParam
}

func (r *Request) RetryCount() int {
	return r.retry
}

func (r *Request) RetryDelay() time.Duration {
	return r.retryTimeout
}

func (r *Request) OnRetry(response *http.Response, err error, retryCount int) {
	respCode := -1
	respData := "<No Data>"
	if response != nil {
		respCode = response.StatusCode
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(response.Body); err != nil {
			respData = "<Read Error>"
		}
		respData = buf.String()
	}
	if r.log != nil {
		r.log.Errorf(
			"%T error %s, the request will be retried %d more times: %d, body: %s\n",
			*r, err, retryCount, respCode, respData,
		)
	}
	return
}

func (r *Request) BuildRaw(ctx context.Context, method string, Url string, authToken string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, Url, body)
	if err != nil {
		return nil, err
	}
	r.addUrlParams(req)
	headers := r.ensureHeaders(req)

	headers.Add("Content-type", "application/json")
	if len(authToken) != 0 {
		headers.Add("Authorization", fmt.Sprint("Bearer ", authToken))
	}
	return req, nil
}

func (r *Request) WithLog(l Logger) {
	r.log = l
}

func (r *Request) WithRetry(retry int, timeout time.Duration) {
	r.retry = retry
	r.retryTimeout = timeout
}

func (r *Request) WithQueryParams(param QueryParam) {
	if r.queryParams == nil {
		r.queryParams = param
	} else {
		for k, v := range param {
			if _, ok := r.queryParams[k]; !ok {
				r.queryParams[k] = []string{}
			}
			r.queryParams[k] = append(r.queryParams[k], v...)
		}
	}
}

func (r *Request) WithHeaders(headers http.Header) {
	if r.headers == nil {
		r.headers = map[string][]string{}
	}
	r.headers = headers
}

func (r *Request) addUrlParams(req *http.Request) {
	rawQuery := req.URL.Query()
	for pn, pv := range r.queryParams {
		for _, v := range pv {
			rawQuery.Add(pn, v)
		}
	}
	req.URL.RawQuery = rawQuery.Encode()
}

func (r *Request) ensureHeaders(req *http.Request) *http.Header {
	if r.headers != nil {
		req.Header = r.headers
	} else {
		req.Header = http.Header{}
	}
	return &req.Header
}

package clo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"io"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type ResponseUnmarshaler func(body io.Reader, dst any) error

type ApiClient struct {
	HttpClient HttpClient
	Log        Logger
	conf       Config
	Unmarshall ResponseUnmarshaler
}

func NewDefaultClient(authKey string, baseUrl string) (*ApiClient, error) {
	return NewDefaultClientFromConfig(Config{AuthKey: authKey, BaseUrl: baseUrl})
}

func NewDefaultClientFromConfig(cfg Config) (*ApiClient, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	cli := http.DefaultClient
	cli.Timeout = cfg.HttpTimeout
	s := &ApiClient{
		HttpClient: http.DefaultClient,
		conf:       cfg,
		Unmarshall: UnmarshallJsonResponse,
	}
	return s, nil
}

func (cli *ApiClient) DoRequest(ctx context.Context, req RequestInt, resp ResponseInterface) (err error) {
	retryCount := req.RetryCount()
	retryDelay := req.RetryDelay()
	if retryCount < 0 {
		return fmt.Errorf("retry number should be positive")
	}

	rawReq, err := req.Build(ctx, cli.conf.BaseUrl, cli.conf.AuthKey)
	if err != nil {
		return err
	}

	var rawResp *http.Response
	for {
		rawResp, err = cli.HttpClient.Do(rawReq)
		if cli.Log != nil {
			cli.Log.Tracef("resp: %v, err: %v\n", rawResp, err)
		}
		if err == nil {
			if rawResp.StatusCode < 500 {
				return cli.parseResponse(rawResp, resp)
			}
			err = cli.parseErrorResponse(rawResp)
		}
		retryCount -= 1
		req.OnRetry(rawResp, err, retryCount)
		if retryCount <= 0 {
			break
		}
		if retryDelay > 0 {
			time.Sleep(retryDelay)
		}

	}
	return err
}

func (cli *ApiClient) parseResponse(resp *http.Response, dst ResponseInterface) error {
	defer resp.Body.Close()
	if request_tools.IsError(resp.StatusCode) {
		return cli.parseErrorResponse(resp)
	}
	if dst != nil {
		err := cli.Unmarshall(resp.Body, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *ApiClient) parseErrorResponse(resp *http.Response) error {
	var de request_tools.DefaultError
	de.Code = resp.StatusCode
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&de)
	switch {
	case err == io.EOF:
		return fmt.Errorf("error with an empty body, status code is : %d", resp.StatusCode)
	case err != nil:
		var b bytes.Buffer
		str := "<FAIL TO READ>"
		if _, err := b.ReadFrom(resp.Body); err != nil {
			str = b.String()
		}
		return fmt.Errorf("can't decode an error body: %s", str)
	}
	return de
}

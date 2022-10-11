package storage

import (
	"github.com/clo-ru/cloapi-go-client/clo"
	"context"
	"fmt"
	"net/http"
)

const (
	s3UserDeleteEndpoint = "/v1/s3_users/%s"
)

type S3UserDeleteRequest struct {
	clo.Request
	UserID string
}

func (r *S3UserDeleteRequest) Make(ctx context.Context, cli *clo.ApiClient) error {
	rawReq, e := r.buildRequest(ctx, cli.Options)
	if e != nil {
		return e
	}
	_, requestError := r.MakeRequest(rawReq, cli)
	if requestError != nil {
		return requestError
	}
	return nil
}

func (r *S3UserDeleteRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(s3UserDeleteEndpoint, r.UserID)
	rawReq, e := http.NewRequestWithContext(
		ctx, http.MethodDelete, baseUrl, nil,
	)
	if e != nil {
		return nil, e
	}
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("Bearer %s", authKey))
	r.WithHeaders(h)
	return rawReq, nil
}

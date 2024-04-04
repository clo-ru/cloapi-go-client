package sshkeys

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	keypairListEndpoint = "%s/v2/projects/%s/keypairs"
)

type KeyPairListRequest struct {
	clo.Request
	ProjectID string
}

func (r *KeyPairListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*KeyPairListResponse, error) {
	resp := &KeyPairListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *KeyPairListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(keypairListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *KeyPairListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *KeyPairListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }

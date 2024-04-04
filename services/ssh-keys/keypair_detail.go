package sshkeys

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	keypairDetailEndpoint = "%s/v2/keypairs/%s/detail"
)

type KeyPairDetailRequest struct {
	clo.Request
	KeypairID string
}

func (r *KeyPairDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*KeyPairDetailResponse, error) {
	resp := &KeyPairDetailResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *KeyPairDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(keypairDetailEndpoint, baseUrl, r.KeypairID), authToken, nil)
}

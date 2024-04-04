package sshkeys

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	keypairDeleteEndpoint = "%s/v2/keypairs/%s"
)

type KeyPairDeleteRequest struct {
	clo.Request
	KeypairID string
}

func (r *KeyPairDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *KeyPairDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(keypairDeleteEndpoint, baseUrl, r.KeypairID), authToken, nil)
}

package sshkeys

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	keypairCreateEndpoint = "%s/v2/projects/%s/keypairs"
)

type KeyPairCreateRequest struct {
	clo.Request
	ProjectID string
	Body      KeyPairCreateBody
}

type KeyPairCreateBody struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

func (r *KeyPairCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	resp := &clo.ResponseCreated{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *KeyPairCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(keypairCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}

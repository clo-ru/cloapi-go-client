package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	addressCreateEndpoint = "%s/v2/projects/%s/addresses"
)

type AddressCreateRequest struct {
	clo.Request
	ProjectID string
	Body      AddressCreateBody
}

type AddressCreateBody struct {
	DdosProtection bool `json:"ddos_protection"`
}

func (r *AddressCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *AddressCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(addressCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}

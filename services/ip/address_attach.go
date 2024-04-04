package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	addressAttachEndpoint = "%s/v2/addresses/%s/attach"
)

type AddressAttachRequest struct {
	clo.Request
	AddressID string
	Body      AddressAttachBody
}

type AddressAttachBody struct {
	ID     string `json:"id"`
	Entity string `json:"entity"`
}

func (r *AddressAttachRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *AddressAttachRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(addressAttachEndpoint, baseUrl, r.AddressID), authToken, body)
}

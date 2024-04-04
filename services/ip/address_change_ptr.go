package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	addressPtrChangeEndpoint = "%s/v2/addresses/%s/ptr"
)

type AddressPtrChangeRequest struct {
	clo.Request
	AddressID string
	Body      AddressPtrChangeBody
}

type AddressPtrChangeBody struct {
	Value string `json:"value"`
}

func (r *AddressPtrChangeRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *AddressPtrChangeRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(addressPtrChangeEndpoint, baseUrl, r.AddressID), authToken, body)
}

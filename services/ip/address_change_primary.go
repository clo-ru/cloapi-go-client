package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	addressPrimaryChangeEndpoint = "%s/v2/addresses/%s/primary"
)

type AddressPrimaryChangeRequest struct {
	clo.Request
	AddressID string
}

func (r *AddressPrimaryChangeRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *AddressPrimaryChangeRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(addressPrimaryChangeEndpoint, baseUrl, r.AddressID), authToken, nil)
}

package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	addressDeleteEndpoint = "%s/v2/addresses/%s"
)

type AddressDeleteRequest struct {
	clo.Request
	AddressID string
}

func (r *AddressDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *AddressDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(addressDeleteEndpoint, baseUrl, r.AddressID), authToken, nil)
}

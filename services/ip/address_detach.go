package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	addressDetachEndpoint = "%s/v2/addresses/%s/detach"
)

type AddressDetachRequest struct {
	clo.Request
	AddressID string
}

func (r *AddressDetachRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *AddressDetachRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(addressDetachEndpoint, baseUrl, r.AddressID), authToken, nil)
}

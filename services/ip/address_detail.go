package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	addressDetailEndpoint = "%s/v2/addresses/%s/detail"
)

type AddressDetailRequest struct {
	clo.Request
	AddressID string
}

func (r *AddressDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*AddressDetailResponse, error) {
	res := &AddressDetailResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *AddressDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(addressDetailEndpoint, baseUrl, r.AddressID), authToken, nil)
}

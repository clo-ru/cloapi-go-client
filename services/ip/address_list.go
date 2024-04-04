package ip

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	addressListEndpoint = "%s/v2/projects/%s/addresses"
)

type AddressListRequest struct {
	clo.Request
	ProjectID string
}

func (r *AddressListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*AddressListResponse, error) {
	resp := &AddressListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}
func (r *AddressListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(addressListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}
func (r *AddressListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *AddressListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }

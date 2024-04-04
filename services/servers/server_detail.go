package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverDetailEndpoint = "%s/v2/servers/%s/detail"
)

type ServerDetailRequest struct {
	clo.Request
	ServerID string
}

func (r *ServerDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*ServerDetailResponse, error) {
	res := &ServerDetailResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *ServerDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(serverDetailEndpoint, baseUrl, r.ServerID), authToken, nil)
}

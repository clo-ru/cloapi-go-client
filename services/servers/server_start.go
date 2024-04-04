package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverStartEndpoint = "%s/v2/servers/%s/start"
)

type ServerStartRequest struct {
	clo.Request
	ServerID string
}

func (r *ServerStartRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerStartRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverStartEndpoint, baseUrl, r.ServerID), authToken, nil)
}

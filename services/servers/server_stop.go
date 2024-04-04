package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverStopEndpoint = "%s/v2/servers/%s/stop"
)

type ServerStopRequest struct {
	clo.Request
	ServerID string
}

func (r *ServerStopRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerStopRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverStopEndpoint, baseUrl, r.ServerID), authToken, nil)
}

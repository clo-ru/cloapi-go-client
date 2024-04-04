package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverRebootEndpoint = "%s/v2/servers/%s/reboot"
)

type ServerRebootRequest struct {
	clo.Request
	ServerID string
}

func (r *ServerRebootRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerRebootRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverRebootEndpoint, baseUrl, r.ServerID), authToken, nil)
}

package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	serverResizeEndpoint = "%s/v2/servers/%s/resize"
)

type ServerResizeRequest struct {
	clo.Request
	ServerID string
	Body     ServerResizeBody
}

type ServerResizeBody struct {
	Ram   int `json:"ram"`
	Vcpus int `json:"vcpus"`
}

func (r *ServerResizeRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerResizeRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverResizeEndpoint, baseUrl, r.ServerID), authToken, body)
}

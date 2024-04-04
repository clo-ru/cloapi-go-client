package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	serverChangePasswdEndpoint = "%s/v2/servers/%s/password"
)

type ServerChangePasswdRequest struct {
	clo.Request
	ServerID string
	Body     ServerChangePasswdBody
}

type ServerChangePasswdBody struct {
	Password string `json:"password"`
}

func (r *ServerChangePasswdRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerChangePasswdRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverChangePasswdEndpoint, baseUrl, r.ServerID), authToken, body)
}

package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	serverDeleteEndpoint = "%s/v2/servers/%s"
)

type ServerDeleteRequest struct {
	clo.Request
	ServerID string
	Body     ServerDeleteBody
}

type ServerDeleteBody struct {
	DeleteVolumes   []string `json:"delete_volumes,omitempty"`
	DeleteAddresses []string `json:"delete_addresses,omitempty"`
	ClearFstab      bool     `json:"clear_fstab"`
}

func (r *ServerDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *ServerDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(serverDeleteEndpoint, baseUrl, r.ServerID), authToken, body)
}

package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverCreateEndpoint = "/v1/projects/%s/servers"
)

type ServerCreateRequest struct {
	clo.Request
	ProjectID string
	Body      ServerCreateBody
}

type ServerCreateBody struct {
	Name      string                `json:"name"`
	Image     string                `json:"image"`
	Recipe    string                `json:"recipe,omitempty"`
	Keypairs  []string              `json:"keypairs,omitempty"`
	Flavor    ServerFlavorBody      `json:"flavor"`
	Storages  []ServerStorageBody   `json:"storages"`
	Licenses  []ServerLicenseBody   `json:"licenses,omitempty"`
	Addresses []ServerAddressesBody `json:"addresses"`
}

type ServerAddressesBody struct {
	External       bool   `json:"external"`
	DdosProtection bool   `json:"ddos_protection"`
	WithFloating   bool   `json:"with_floating,omitempty"`
	FloatingIpID   string `json:"floatingip_id,omitempty"`
	Version        int    `json:"version"`
}

type ServerFlavorBody struct {
	Ram   int `json:"ram"`
	Vcpus int `json:"vcpus"`
}

type ServerLicenseBody struct {
	Value int    `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
	Addon string `json:"addon,omitempty"`
}

type ServerStorageBody struct {
	Size        int    `json:"size"`
	Bootable    bool   `json:"bootable"`
	StorageType string `json:"storage_type"`
}

func (r *ServerCreateRequest) Make(ctx context.Context, cli *clo.ApiClient) (ServerCreateResponse, error) {
	rawReq, e := r.buildRequest(ctx, cli.Options)
	if e != nil {
		return ServerCreateResponse{}, e
	}
	rawResp, requestError := r.MakeRequest(rawReq, cli)
	if requestError != nil {
		return ServerCreateResponse{}, requestError
	}
	defer rawResp.Body.Close()
	var resp ServerCreateResponse
	if e = resp.FromJsonBody(rawResp.Body); e != nil {
		return ServerCreateResponse{}, e
	}
	return resp, nil
}

func (r *ServerCreateRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(serverCreateEndpoint, r.ProjectID)
	b := new(bytes.Buffer)
	if e := json.NewEncoder(b).Encode(r.Body); e != nil {
		return nil, fmt.Errorf("can't encode body parameters, %s", e.Error())
	}
	rawReq, e := http.NewRequestWithContext(
		ctx, http.MethodPost, baseUrl, b,
	)
	if e != nil {
		return nil, e
	}
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("Bearer %s", authKey))
	r.WithHeaders(h)
	return rawReq, nil
}

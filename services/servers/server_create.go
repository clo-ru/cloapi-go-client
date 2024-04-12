package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	serverCreateEndpoint = "%s/v2/projects/%s/servers"
)

type ServerCreateRequest struct {
	clo.Request
	ProjectID string
	Body      ServerCreateBody
}

type ServerCreateBody struct {
	Name         string                `json:"name"`
	Image        string                `json:"image,omitempty"`
	Recipe       string                `json:"recipe,omitempty"`
	SourceVolume string                `json:"volume,omitempty"`
	Flavor       ServerFlavorBody      `json:"flavor"`
	Storages     []ServerStorageBody   `json:"storages,omitempty"`
	Licenses     []ServerLicenseBody   `json:"licenses,omitempty"`
	Addresses    []ServerAddressesBody `json:"addresses,omitempty"`
	Keypairs     []string              `json:"keypairs,omitempty"`
}

type ServerAddressesBody struct {
	External       bool   `json:"external"`
	DdosProtection bool   `json:"ddos_protection,omitempty"`
	AddressId      string `json:"address_id,omitempty"`
	MaxBandwidth   int    `json:"bandwidth_max_mbps,omitempty"`
	Version        int    `json:"version,omitempty"`
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
	StorageType string `json:"storage_type,omitempty"`
}

func (r *ServerCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *ServerCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(serverCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}

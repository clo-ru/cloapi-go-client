package disks

import (
	"bytes"
	"github.com/clo-ru/cloapi-go-client/clo"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	volumeDetachEndpoint = "/v1/volumes/%s/detach"
)

type VolumeDetachRequest struct {
	clo.Request
	VolumeID string
	Body     VolumeDetachBody
}

type VolumeDetachBody struct {
	Force bool `json:"force"`
}

func (r *VolumeDetachRequest) Make(ctx context.Context, cli *clo.ApiClient) error {
	rawReq, e := r.buildRequest(ctx, cli.Options)
	if e != nil {
		return e
	}
	_, requestError := r.MakeRequest(rawReq, cli)
	if requestError != nil {
		return requestError
	}
	return nil

}

func (r *VolumeDetachRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(volumeDetachEndpoint, r.VolumeID)
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

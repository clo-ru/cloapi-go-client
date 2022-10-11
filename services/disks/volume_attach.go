package disks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	volumeAttachEndpoint = "/v1/volumes/%s/attach"
)

type VolumeAttachRequest struct {
	clo.Request
	VolumeID string
	Body     VolumeAttachBody
}

type VolumeAttachBody struct {
	MountPath string `json:"mount_path,omitempty"`
	ServerID  string `json:"server_id"`
}

func (r *VolumeAttachRequest) Make(ctx context.Context, cli *clo.ApiClient) (VolumeAttachResponse, error) {
	rawReq, e := r.buildRequest(ctx, cli.Options)
	if e != nil {
		return VolumeAttachResponse{}, e
	}
	rawResp, requestError := r.MakeRequest(rawReq, cli)
	if requestError != nil {
		return VolumeAttachResponse{}, requestError
	}
	defer rawResp.Body.Close()
	var resp VolumeAttachResponse
	if e = resp.FromJsonBody(rawResp.Body); e != nil {
		return VolumeAttachResponse{}, e
	}
	return resp, nil
}

func (r *VolumeAttachRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(volumeAttachEndpoint, r.VolumeID)
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

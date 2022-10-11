package ip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	fipPtrChangeEndpoint = "/v1/floatingips/%s/ptr"
)

type FipPtrChangeRequest struct {
	clo.Request
	FloatingID string
	Body       FipPtrChangeBody
}

type FipPtrChangeBody struct {
	Value string `json:"value"`
}

func (r *FipPtrChangeRequest) Make(ctx context.Context, cli *clo.ApiClient) error {
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

func (r *FipPtrChangeRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(fipPtrChangeEndpoint, r.FloatingID)
	b := new(bytes.Buffer)
	if e := json.NewEncoder(b).Encode(r.Body); e != nil {
		return nil, fmt.Errorf("can't encode body parameters, %s", e.Error())
	}
	rawReq, e := http.NewRequestWithContext(
		ctx, http.MethodPut, baseUrl, b,
	)
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("Bearer %s", authKey))
	r.WithHeaders(h)
	if e != nil {
		return nil, e
	}
	return rawReq, nil
}

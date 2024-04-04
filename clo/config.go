package clo

import (
	"encoding/json"
	"fmt"
	"time"
)

type Config struct {
	AuthKey     string        `json:"auth_key"`
	BaseUrl     string        `json:"base_url"`
	HttpTimeout time.Duration `json:"http_timeout"`
}

func (cfg Config) FromMap(opt map[string]interface{}) error {
	b, e := json.Marshal(opt)
	if e != nil {
		return e
	}
	if e := json.Unmarshal(b, &cfg); e != nil {
		return e
	}
	return nil
}

func (cfg Config) Validate() error {
	if len(cfg.BaseUrl) == 0 {
		return fmt.Errorf("BaseUrl should be provided")
	}
	if len(cfg.AuthKey) == 0 {
		return fmt.Errorf("AuthKey should be provided")
	}
	return nil
}

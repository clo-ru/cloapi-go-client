package sshkeys

import "github.com/clo-ru/cloapi-go-client/clo"

type KeyPair struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

type KeyPairDetailResponse = clo.Response[KeyPair]
type KeyPairListResponse = clo.ListResponse[KeyPair]

package ip

import "github.com/clo-ru/cloapi-go-client/clo"

type Address struct {
	ID             string             `json:"id"`
	MacAddr        string             `json:"mac_addr"`
	Mask           string             `json:"mask"`
	Ptr            string             `json:"ptr"`
	Type           string             `json:"type"`
	Status         string             `json:"status"`
	Version        int                `json:"version"`
	CreatedIn      string             `json:"created_in"`
	UpdatedIn      string             `json:"updated_in"`
	Address        string             `json:"address"`
	DdosProtection bool               `json:"ddos_protection"`
	External       bool               `json:"external"`
	IsPrimary      bool               `json:"is_primary"`
	Bandwidth      int                `json:"bandwidth_max_mbps"`
	AttachedTo     *AttachedToDetails `json:"attached_to"`
}

type AttachedToDetails struct {
	ID     string `json:"id"`
	Entity string `json:"entity"`
}

type AddressDetailResponse = clo.Response[Address]
type AddressListResponse = clo.ListResponse[Address]

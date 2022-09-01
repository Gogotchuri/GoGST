package vayanaTypes

import "strings"

type DataResponse struct {
	Data string `json:"data"`
}

func (d *DataResponse) GetData() []byte {
	d.Data = strings.ReplaceAll(d.Data, "\\", "")
	d.Data = strings.ReplaceAll(d.Data, "\"{", "{")
	d.Data = strings.ReplaceAll(d.Data, "}\"", "}")
	return []byte(d.Data)
}

type GSTINDetails struct {
	Gstin     string `json:"gstin"`
	TradeName string `json:"tradeName"`
	LegalName string `json:"legalName"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	StateCode string `json:"stateCode"`
	PinCode   string `json:"pinCode"`
	TxpType   string `json:"txpType"`
	Status    string `json:"status"`
	BlkStatus string `json:"blkStatus"`
}

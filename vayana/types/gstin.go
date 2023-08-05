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

type Address struct {
	StateName      string `json:"stcd"`
	Location       string `json:"loc"`
	PINCode        string `json:"pncd"`
	StreetName     string `json:"st"`
	BuildingNumber string `json:"bno"`
	BuildingName   string `json:"bnm"`
	FloorNumber    string `json:"flno"`
	Latitude       string `json:"lt"`
	Longitude      string `json:"lg"`
}

type AddressAndNature struct {
	Address Address  `json:"addr"`
	Nature  []string `json:"ntr"`
}

type GSTINPublicDetails struct {
	AdditionalAddresses      []AddressAndNature `json:"adadr"`
	PrimaryAddress           AddressAndNature   `json:"pradr"`
	ConstitutionOfBusiness   string             `json:"ctb"`
	CenterOfJurisdiction     string             `json:"ctj"`
	StateOfJurisdictionCode  string             `json:"stjCd"`
	DateOfCancellation       string             `json:"cxdt"`
	TaxPayerType             string             `json:"dty"`
	Gstin                    string             `json:"gstin"`
	EInvoiceStatus           string             `json:"einvoiceStatus"`
	LegalName                string             `json:"lgnm"`
	BusinessActivityNature   []string           `json:"nba"`
	LstUpdateDate            string             `json:"rgdt"`
	StateOfJurisdiction      string             `json:"stj"`
	CenterOfJurisdictionCode string             `json:"ctjCd"`
	LastUpdate               string             `json:"lstupdt"`
	GSTINStatus              string             `json:"sts"`
	TradeName                string             `json:"tradeNam"`
}

package vayanaTypes

import (
	"encoding/json"
	"strings"
)

type DataResponse struct {
	Data any `json:"data"`
}

func (d *DataResponse) GetData() ([]byte, error) {
	sData, ok := d.Data.(string)
	if !ok {
		return json.Marshal(d.Data)
	}
	d.Data = strings.ReplaceAll(sData, "\\", "")
	d.Data = strings.ReplaceAll(sData, "\"{", "{")
	d.Data = strings.ReplaceAll(sData, "}\"", "}")
	return []byte(sData), nil
}

type GSTINDetails struct {
	ConstitutionOfBusiness    string          `json:"ctb"`
	CentreJurisdiction        string          `json:"ctj"`
	CentreJurisdictionCode    string          `json:"ctjCd"`
	DateOfCancellation        string          `json:"cxdt"`
	TaxpayerType              string          `json:"dty"`
	EinvoiceStatus            string          `json:"einvoiceStatus"`
	GSTIN                     string          `json:"gstin"`
	LegalNameOfBusiness       string          `json:"lgnm"`
	LastUpdatedDate           string          `json:"lstupddt"`
	NatureOfBusiness          []string        `json:"nba"`
	DateOfRegistration        string          `json:"rgdt"`
	StateJurisdiction         string          `json:"stj"`
	StateJurisdictionCode     string          `json:"stjCd"`
	GSTNStatus                string          `json:"sts"`
	TradeName                 string          `json:"tradeNam"`
	AdditionalPlaceOfBusiness []BusinessPlace `json:"adadr"`
	PrincipalPlaceOfBusiness  BusinessPlace   `json:"pradr"`
}

type Address struct {
	BuildingName string `json:"bnm"`
	DoorNumber   string `json:"bno"`
	District     string `json:"dst"`
	FloorNumber  string `json:"flno"`
	GeoCodeLevel string `json:"geocodelvl"`
	Landmark     string `json:"landMark"`
	Longitude    string `json:"lg"`
	Location     string `json:"loc"`
	Locality     string `json:"locality"`
	Latitude     string `json:"lt"`
	PinCode      string `json:"pncd"`
	Street       string `json:"st"`
	StateName    string `json:"stcd"`
}

type BusinessPlace struct {
	Addr Address `json:"addr"`
	Ntr  string  `json:"ntr"`
}

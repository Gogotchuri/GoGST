package types

type EWBCreateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Alert   string `json:"alert"`
	Info    string `json:"info"` //Set when status is 0

	Uuid         string `json:"uuid"`
	ValidUpto    string `json:"validUpto"`
	EwayBillNo   uint64 `json:"ewayBillNo,number"`
	EwayBillDate string `json:"ewayBillDate"`
}

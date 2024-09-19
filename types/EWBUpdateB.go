package types

const UpdateBPath = "/update-partb/"

// Update Part B
type EWBUpdateBRequest struct {
	EwayBillNo   EWBNumber `json:"EwbNo"`
	FromPlace    string    `json:"FromPlace"`
	FromState    int       `json:"FromState"`
	ReasonCode   string    `json:"ReasonCode"`
	ReasonRem    string    `json:"ReasonRem"`
	TransDocDate string    `json:"TransDocDate "`
	TransDocNo   string    `json:"TransDocNo "`
	TransMode    string    `json:"TransMode"`
	VehicleNo    string    `json:"VehicleNo"`
}

type EWBUpdateBResponse struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	VehUpdDate string `json:"vehUpdDate"`
	ValidUpto  string `json:"validUpto"`
}

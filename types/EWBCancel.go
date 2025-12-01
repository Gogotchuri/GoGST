package types

const CancelPath = "/cancel/"

type EWBCancelRequest struct {
	CancelRemark     string    `json:"cancelRmrk"`
	CancelReasonCode int       `json:"cancelRsnCode"`
	EwayBillNo       EWBNumber `json:"ewbNo"`
}

// EWBCancelResponseData represents the data object in v3.0 cancel response
type EWBCancelResponseData struct {
	EwayBillNo string `json:"ewayBillNo"`
	CancelDate string `json:"cancelDate"`
}

// EWBCancelResponseV3 represents the v3.0 API response structure for Cancel EWB
type EWBCancelResponseV3 struct {
	Status         string                 `json:"status"`
	Data           EWBCancelResponseData  `json:"data"`
	Error          interface{}            `json:"error"`
	Info           interface{}            `json:"info"`
	AdditionalInfo interface{}            `json:"additionalInfo"`
	Alert          interface{}            `json:"alert"`
}

// EWBCancelResponse represents the legacy v1.0 API response structure (deprecated)
type EWBCancelResponse struct {
	Status     int       `json:"status"`
	Message    string    `json:"message"`
	CancelDate string    `json:"cancelDate"`
	EwayBillNo EWBNumber `json:"ewayBillNo"`
}

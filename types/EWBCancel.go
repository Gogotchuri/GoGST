package types

const CancelPath = "/cancel/"

type EWBCancelRequest struct {
	CancelRemark     string    `json:"cancelRmrk"`
	CancelReasonCode int       `json:"cancelRsnCode"`
	EwayBillNo       EWBNumber `json:"ewbNo"`
}

type EWBCancelResponse struct {
	Status     int       `json:"status"`
	Message    string    `json:"message"`
	CancelDate string    `json:"cancelDate"`
	EwayBillNo EWBNumber `json:"ewayBillNo"`
}

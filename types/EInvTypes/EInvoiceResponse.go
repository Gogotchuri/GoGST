package EInvTypes

import "github.com/gogotchuri/GoGST/types"



type Response struct {
	AcknowledgementNo types.AmbiguousString          `json:"AckNo"`
	AcknowledgementDt string          `json:"AckDt"`
	IRN               string          `json:"Irn"`
	SignedInvoice     string          `json:"SignedInvoice"`
	SignedQRCode      string          `json:"SignedQRCode"`
	Status            string          `json:"Status"`
	EwbNo             types.EWBNumber `json:"EwbNo"`
	EwbDate           string          `json:"EwbDt"`
	EwbValidTill      string          `json:"EwbValidTill"`
	Remarks           string          `json:"Remarks"`
}

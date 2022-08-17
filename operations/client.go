package operations

import "github.com/gogotchuri/GoGST/types"

type Client interface {
	CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error)
	CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error)

	GetGSTIN(gstin string) (*types.GSTINResponse, error)
}

var _ Client = &client{}

type client struct {
	token   string
	baseURL string
	gstin   string
}

func NewClient(token, baseURL, gstin string) *client {
	return &client{
		token:   token,
		baseURL: baseURL,
		gstin:   gstin,
	}
}

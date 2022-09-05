package GoGST

import (
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
)

type Client interface {
	Ping() error
	Authenticate(email, password string) error
	IsAuthenticated() (bool, error)
	SetActiveToken(token string)
	Logout() error

	CreateGSPClient(gstin, username, password string) (GSPClient, error)
}

type GSPClient interface {
	CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error)
	GetEWayBill(ewbNo string) (*types.EWBGetResponse, error)
	CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error)
	GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetails, error)
}

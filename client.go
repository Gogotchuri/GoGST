package GoGST

import (
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
)

type Client interface {
	Authenticate(email, password string) error
	SetActiveToken(token string)
	Logout() error

	CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error)
	GetEWayBill(ewbNo string) (*types.EWBGetResponse, error)
	CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error)

	GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetails, error)
	Ping() error
}

package GoGST

import (
	"github.com/gogotchuri/GoGST/types"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"strings"
	"time"
)

type Client interface {
	Ping() error
	Authenticate(email, password string) error
	IsAuthenticated() (bool, error)
	SetActiveToken(token string)
	Logout() error

	CreateGSPClient(gstin, username, password string) (GSPClient, error)
	CreateGSPEInvoicesClient(gstin, username, password string) (GSPEInvoiceClient, error)
}

type GSPClient interface {
	CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error)
	GetEWayBill(ewbNo string) (*types.EWBGetResponse, error)
	CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error)
	GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetails, error)
	GetEWayBillsByDate(date time.Time) ([]types.EWBGetResponse, error)
}

type GSPEInvoiceClient interface {
	CreateEInvoice(eInv EInvTypes.EInvoiceCreate) (*EInvTypes.Response, error)
	GetEInvoice(irn string) (*EInvTypes.Response, error)
}

func IsGSPCredentialsError(err error) bool {
	return strings.Contains(err.Error(), "Invalid login credentialsa")
}

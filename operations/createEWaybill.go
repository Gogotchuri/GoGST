package operations

import (
	"github.com/gogotchuri/GoGST/types"
	"net/http"
)

const (
	token   = " bfda565b-fe89-4300-8150-5f86dee7d060"
	baseURL = "https://my.gstzen.in/~gstzen/a/ewbapi/generate"
)

func CreateEWaybill(ewb *types.EWBCreateRequest) (string, error) {
	headers := http.Header{
		"Content-Type": {"application/json"},
		"Token":        {token},
		"gstin":        {ewb.FromGstin},
	}
	body := ewb.ToJSON()
	req, err := http.NewRequest(http.MethodPost, baseURL)
	if err != nil {
		return "", err
	}
	req.Header = headers
	return "", nil
}

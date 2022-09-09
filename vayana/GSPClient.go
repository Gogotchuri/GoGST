package vayana

import (
	"fmt"
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
	"strings"
	"time"
)

const vayanaBasicEWBBase = "/basic/ewb/v1.0/v1.03"

var _ GoGST.GSPClient = &gspClient{}

type gspClient struct {
	theodoreClient *client
	httpClient     *http.Client

	creatorGSTIN string
	username     string
	password     string
}

func (c *gspClient) GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetails, error) {
	endpoint := fmt.Sprintf("/basic/ewb/v1.0/v1.03/gstin-details/%s", gstin)
	resp := &vayanaTypes.GSTINDetails{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) GetEWayBillsByDate(date time.Time) ([]types.EWBGetResponse, error) {
	endpoint := vayanaBasicEWBBase + fmt.Sprintf("/by-date/%d/%d/%d", date.Day(), date.Month(), date.Year())
	resp := make([]types.EWBGetResponse, 0)
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     &resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create ewaybill: %s", err)
	}
	return resp, nil
}
func (c *gspClient) CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error) {
	endpoint := vayanaBasicEWBBase + "/gen-ewb"
	resp := &types.EWBCreateResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     ewb,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create ewaybill: %s", err)
	}
	return resp, nil
}

func (c *gspClient) CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error) {
	endpoint := fmt.Sprintf("%s/cancel", vayanaBasicEWBBase)
	resp := &types.EWBCancelResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     cancel,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) GetEWayBill(ewbNo string) (*types.EWBGetResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", vayanaBasicEWBBase, ewbNo)
	resp := &types.EWBGetResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) handleError(vErr *vayanaTypes.Error) error {
	if vErr == nil {
		return nil
	}
	if vErr.IsEWBError() {
		messages := strings.Join(vErr.GetErrorMessages(), ";")
		return fmt.Errorf("%s", messages)
	} else if vErr.IsTokenExpired() {
		return vayanaTypes.ErrorTokenExpired
	}
	return nil
}

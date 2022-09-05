package vayana

import (
	"fmt"
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
	"strings"
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
	err, _ := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		return nil, err
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
		if vErr != nil {
			if vErr.IsEWBError() {
				messages := strings.Join(vErr.GetErrorMessages(), "; ")
				err = fmt.Errorf("%s", messages)
			} else if vErr.IsTokenExpired() {
				return resp, vayanaTypes.ErrorTokenExpired
			}
		}
		return resp, fmt.Errorf("failed to create ewaybill: %s", err)
	}
	return resp, nil
}

func (c *gspClient) CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error) {
	endpoint := fmt.Sprintf("%s/cancel", vayanaBasicEWBBase)
	resp := &types.EWBCancelResponse{}
	err, _ := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     cancel,
		dest:     resp,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) GetEWayBill(ewbNo string) (*types.EWBGetResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", vayanaBasicEWBBase, ewbNo)
	resp := &types.EWBGetResponse{}
	err, _ := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

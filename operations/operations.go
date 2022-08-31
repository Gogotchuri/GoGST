package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogotchuri/GoGST/types"
	"io/ioutil"
	"net/http"
)

const (
	token                = "43af8d80-5da7-4479-9b55-7fe342c9eb62"
	baseURL              = "https://my.gstzen.in/~gstzen/a/ewbapi/generate/"
	vayanaBaseURL        = "https://solo.enriched-api.vayana.com"
	vayanaOrganizationID = "3b5f5dfa-d244-40ed-bfb1-04d7820328b5" //X-FLYNN-N-ORG-ID header
)

func (c client) CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error) {
	var ewbRes types.EWBCreateResponse
	if err := c.makeGSTZRequest(types.CreatePath, ewb, &ewbRes, c.gstin); err != nil {
		return nil, err
	}
	if ewbRes.Status == 0 {
		return nil, fmt.Errorf("request failed: %s", ewbRes.Message)
	}
	return &ewbRes, nil
}

func (c client) CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error) {
	var cancelRes types.EWBCancelResponse
	if err := c.makeGSTZRequest(types.CancelPath, cancel, &cancelRes, c.gstin); err != nil {
		return nil, err
	}
	if cancelRes.Status == 0 {
		return nil, fmt.Errorf("request failed: %s", cancelRes.Message)
	}
	return &cancelRes, nil
}

func (c client) GetGSTIN(gstin string) (*types.GSTINResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c client) makeGSTZRequest(path, req, res interface{}, gstin string) error {
	encoded, err := json.Marshal(req)
	request, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(encoded))
	if err != nil {
		return fmt.Errorf("can't create request: %v", err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)
	request.Header.Add("gstin", gstin)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("can't send request: %v", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode >= 300 {
		return fmt.Errorf("request failed with status %s error: %s", response.Status, string(body))
	}
	if err := json.Unmarshal(body, res); err != nil {
		return fmt.Errorf("can't unmarshal response: %v", err)
	}
	return nil
}

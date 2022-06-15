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
	token   = "43af8d80-5da7-4479-9b55-7fe342c9eb62"
	baseURL = "https://my.gstzen.in/~gstzen/a/ewbapi/generate/"
)

func (c client) CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error) {
	encoded, err := json.Marshal(ewb)
	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(encoded))
	if err != nil {
		return nil, fmt.Errorf("can't create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)
	req.Header.Add("gstin", ewb.FromGstin)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't send request: %v", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status %s error: %s", response.Status, string(body))
	}
	var ewbRes types.EWBCreateResponse
	if err := json.Unmarshal(body, &ewbRes); err != nil {
		return nil, fmt.Errorf("can't unmarshal response: %v", err)
	}
	if ewbRes.Status == 0 {
		return nil, fmt.Errorf("request failed: %s", ewbRes.Message)
	}
	return &ewbRes, nil
}

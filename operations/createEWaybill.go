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

func CreateEWaybill(ewb types.EWBCreateRequest) ([]byte, error) {
	encoded, err := json.Marshal(ewb)
	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token)
	req.Header.Add("gstin", ewb.FromGstin)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("response Body:", string(body))
	return body, nil
}

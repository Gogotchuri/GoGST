package vayana

import (
	"encoding/json"
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
)

func (c *gspClient) sendRequest(r request) (error, *vayanaTypes.Error) {
	if ok, err := c.theodoreClient.IsAuthenticated(); !ok {
		return fmt.Errorf("token is empty, athenticate first. %s", err.Error()), nil
	}
	req, err := c.theodoreClient.newRequest(r.method, c.theodoreClient.getEndpointURL(r.endpoint, false), r.body, false)
	if err != nil {
		return err, nil
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.theodoreClient.token))
	req.Header.Set("X-FLYNN-N-USER-TOKEN", c.theodoreClient.token)
	req.Header.Set("X-FLYNN-N-ORG-ID", c.theodoreClient.organizationID)

	req.Header.Set("X-FLYNN-N-IRP-GSTIN", c.creatorGSTIN)
	req.Header.Set("X-FLYNN-N-IRP-USERNAME", c.username)
	req.Header.Set("X-FLYNN-N-IRP-PWD", c.password)
	req.Header.Set("X-FLYNN-N-IRP-GSP-CODE", "clayfin")

	req.Header.Set("X-FLYNN-N-EWB-GSTIN", c.creatorGSTIN)
	req.Header.Set("X-FLYNN-N-EWB-USERNAME", c.username)
	req.Header.Set("X-FLYNN-N-EWB-PWD", c.password)
	req.Header.Set("X-FLYNN-N-EWB-GSP-CODE", "clayfin")
	destRaw := &vayanaTypes.DataResponse{}
	err, vErr := c.theodoreClient.send(req, destRaw, false)
	if err != nil {
		return err, vErr
	}
	err = json.Unmarshal(destRaw.GetData(), r.dest)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

package vayana

import (
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
)

func (c *client) GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetailsResponse, error) {
	endpoint := fmt.Sprintf("/basic/einv/v1.0/nic/eivital/v1.03/Master/gstin/%s", gstin)
	resp := &vayanaTypes.GSTINDetailsResponse{}
	err := c.makeAuthorizedRequest(http.MethodGet, endpoint, nil, resp, false, false)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

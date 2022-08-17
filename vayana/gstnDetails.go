package vayana

import (
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
)

func (c *client) GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetailsResponse, error) {
	endpoint := fmt.Sprintf("/basic/einv/1.0/nic/eivital/1.03/master/gstin/%s", gstin)
	resp := &vayanaTypes.GSTINDetailsResponse{}
	err := c.makeAuthorizedRequest(http.MethodGet, endpoint, nil, resp, true, false)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

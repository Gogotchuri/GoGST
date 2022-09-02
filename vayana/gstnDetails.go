package vayana

import (
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
)

func (c *client) GetGSTINDetails(gstin string) (*vayanaTypes.GSTINDetails, error) {
	endpoint := fmt.Sprintf("/basic/ewb/v1.0/v1.03/gstin-details/%s", gstin)
	resp := &vayanaTypes.GSTINDetails{}
	err, _ := c.sendAuthorizedRequest(http.MethodGet, endpoint, nil, resp, false, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

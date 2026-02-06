package vayana

import (
	"fmt"
	"net/http"

	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
)

var _ GoGST.GSPEInvoiceClient = &gspClient{}

const (
	// v3.0 API - Current recommended version
	vayanaEInvVersion    = "v3.0"
	vayanaEInvAPIVersion = "v1.03"
	// EInv provider - ni1 replaces deprecated nic
	vayanaEInvProvider = "ni1"
)

// getEInvBase returns the base path for E-Invoice v3.0 APIs
func (c *gspClient) getEInvBase() string {
	return fmt.Sprintf("/basic/einv/%s/%s/eicore/%s/Invoice", vayanaEInvVersion, vayanaEInvProvider, vayanaEInvAPIVersion)
}

func (c *gspClient) CreateEInvoice(eInv EInvTypes.EInvoiceCreate) (*EInvTypes.Response, error) {
	c.validationLock.Lock()
	validationError := eInv.Validate(c.validator)
	c.validationLock.Unlock()
	if validationError != nil {
		return nil, validationError
	}
	endpoint := c.getEInvBase()
	resp := &EInvTypes.Response{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     eInv,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create e-invoice: %s", err)
	}
	return resp, nil
}

func (c *gspClient) GetEInvoice(irn string) (*EInvTypes.Response, error) {
	endpoint := fmt.Sprintf("%s/irn/%s", c.getEInvBase(), irn)
	resp := &EInvTypes.Response{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create e-invoice: %s", err)
	}
	return resp, nil
}

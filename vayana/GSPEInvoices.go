package vayana

import (
	"fmt"
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
	"net/http"
)

var _ GoGST.GSPEInvoiceClient = &gspClient{}

const vayanaBasicEInvBase = "/basic/einv/v1.0/nic/eicore/v1.03/Invoice"

func (c *gspClient) CreateEInvoice(eInv EInvTypes.EInvoiceCreate) (*EInvTypes.Response, error) {
	c.validationLock.Lock()
	validationError := eInv.Validate(c.validator)
	c.validationLock.Unlock()
	if validationError != nil {
		return nil, validationError
	}
	endpoint := vayanaBasicEInvBase
	resp := &EInvTypes.Response{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     eInv,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create e-invoice: %s", err)
	}
	return resp, nil
}

func (c *gspClient) GetEInvoice(irn string) (*EInvTypes.Response, error) {
	endpoint := fmt.Sprintf("%s/irn/%s", vayanaBasicEInvBase, irn)
	resp := &EInvTypes.Response{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create e-invoice: %s", err)
	}
	return resp, nil
}

package vayana

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogotchuri/GoGST"
	ewbConsts "github.com/gogotchuri/GoGST/consts"
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"github.com/gogotchuri/go-validator"
)

const (
	// v3.0 API - Current recommended version
	vayanaEWBVersion    = "v3.0"
	vayanaEWBAPIVersion = "v1.03"
	// EWB provider - always use ew1
	vayanaEWBProvider = "ew1"
)

var _ GoGST.GSPClient = &gspClient{}

type gspClient struct {
	validator      *validator.Validate
	validationLock *sync.Mutex

	theodoreClient *client

	creatorGSTIN string
	username     string
	password     string
}

// getEWBBase returns the base path for EWB v3.0 APIs
func (c *gspClient) getEWBBase() string {
	return fmt.Sprintf("/basic/eway/%s/%s/%s/ewayapi", vayanaEWBVersion, vayanaEWBProvider, vayanaEWBAPIVersion)
}

func (c *gspClient) GetTaxPayerDetails(gstin string) (*vayanaTypes.GSTINDetails, error) {
	endpoint := fmt.Sprintf("/basic/gstn/v3.0/commonapi/v1.3/search?gstin=%s&action=TP", gstin)
	resp := &vayanaTypes.GSTINDetails{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) GetEWayBillsByDate(date time.Time) ([]types.EWBGetResponse, error) {
	endpoint := c.getEWBBase() + fmt.Sprintf("/by-date/%d/%d/%d", date.Day(), date.Month(), date.Year())
	resp := make([]types.EWBGetResponse, 0)
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     &resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to get ewaybills by date: %s", err)
	}
	return resp, nil
}

func (c *gspClient) CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error) {
	endpoint := c.getEWBBase() + "/gen-ewb"
	resp := &types.EWBCreateResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     ewb,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return resp, fmt.Errorf("failed to create ewaybill: %s", err)
	}
	return resp, nil
}

func (c *gspClient) CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error) {
	endpoint := c.getEWBBase() + "/cancel"
	resp := &types.EWBCancelResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: endpoint,
		body:     cancel,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) GetEWayBill(ewbNo string) (*types.EWBGetResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", c.getEWBBase(), ewbNo)
	resp := &types.EWBGetResponse{}
	err, vErr := c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: endpoint,
		dest:     resp,
	})
	if err != nil {
		if nErr := c.handleError(vErr, err); nErr != nil {
			return nil, nErr
		}
		return nil, err
	}
	return resp, nil
}

func (c *gspClient) handleError(vErr *vayanaTypes.Error, err error) error {
	// Check vErr first (existing logic)
	if vErr != nil {
		if vErr.IsEWBError() {
			messages := strings.Join(vErr.GetEwbErrorMessages(), ";")
			if messages != "" {
				return fmt.Errorf("%s", messages)
			}
		} else if vErr.IsTokenExpired() {
			return vayanaTypes.ErrorTokenExpired
		} else if vErr.IsIRPError() || vErr.IsInvalidBodyError() {
			return vErr
		}
	}

	// Parse error string for embedded JSON error codes
	if err != nil {
		errStr := err.Error()
		// Find the JSON part (starts with {)
		jsonStart := strings.Index(errStr, "{")
		if jsonStart >= 0 {
			jsonStr := errStr[jsonStart:]
			// Parse the JSON error response
			var errorResponse struct {
				Status string `json:"status"`
				Error  struct {
					Message string `json:"message"`
					Type    string `json:"type"`
					Args    struct {
						Details []struct {
							ErrorCode    string  `json:"ErrorCode"`
							ErrorMessage *string `json:"ErrorMessage"`
						} `json:"details"`
					} `json:"args"`
				} `json:"error"`
				Info string `json:"info"`
			}

			if jsonErr := json.Unmarshal([]byte(jsonStr), &errorResponse); jsonErr == nil {
				// Check if error has error codes in details
				if len(errorResponse.Error.Args.Details) > 0 {
					var messages []string
					for _, detail := range errorResponse.Error.Args.Details {
						code, parseErr := strconv.Atoi(detail.ErrorCode)
						if parseErr == nil {
							if msg, ok := ewbConsts.ErrorCodes[code]; ok {
								messages = append(messages, fmt.Sprintf("[%s] %s", detail.ErrorCode, msg))
							} else {
								messages = append(messages, fmt.Sprintf("[%s] Unknown error code", detail.ErrorCode))
							}
						}
					}
					if len(messages) > 0 {
						result := fmt.Sprintf("GST Errors: %s", strings.Join(messages, "; "))
						if errorResponse.Info != "" {
							result += fmt.Sprintf(" | Info: %s", strings.TrimPrefix(errorResponse.Info, ", "))
						}
						return fmt.Errorf("%s", result)
					}
				}
			}
		}
		return err
	}

	return nil
}

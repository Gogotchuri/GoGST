package vayana

import (
	"fmt"
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	"github.com/gogotchuri/GoGST/vayana/encription"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
	"strings"
	"sync"
	"time"
)

var _ GoGST.Client = &client{}

type client struct {
	apiBaseURL      string
	theodoreBaseURL string
	apiVersion      string
	organizationID  string

	publicKey string
	rek       string

	token          string
	tokenExpiresAt time.Time
	tokenLock      sync.Mutex

	httpClient *http.Client
}

func NewDefaultClient(production bool, organizationID string) (GoGST.Client, error) {
	rek, err := encription.GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate rek: %s", err)
	}
	if production {
		return NewClient(vayanaTypes.BaseProductionAPI, vayanaTypes.TheodoreBaseProductionAPI, vayanaTypes.ProductionPublicKey, organizationID, rek, vayanaTypes.DefaultVersion), nil
	}
	return NewClient(vayanaTypes.BaseSandboxAPI, vayanaTypes.TheodoreBaseSandboxAPI, organizationID, vayanaTypes.SandboxPublicKey, rek, vayanaTypes.DefaultVersion), nil
}

func NewClient(baseURL, theodoreBaseURL, organizationID, publicKey, rek, version string) GoGST.Client {
	return &client{
		apiBaseURL:      baseURL,
		theodoreBaseURL: theodoreBaseURL,
		apiVersion:      version,

		publicKey: publicKey,
		rek:       rek,

		organizationID: organizationID,
		httpClient:     &http.Client{},
	}
}

func (c *client) Ping() error {
	err, _ := c.sendRequest(http.MethodGet, vayanaTypes.HealthCheck, nil, nil, false, true)
	return err
}

func (c *client) SetActiveToken(token string) {
	c.tokenLock.Lock()
	c.token = token
	c.tokenExpiresAt = time.Now().Add(359 * time.Minute)
	c.tokenLock.Unlock()
}

func (c *client) Authenticate(email, password string) error {
	resp := vayanaTypes.AuthResponse{}
	err, _ := c.sendRequest(http.MethodPost, vayanaTypes.AuthTokens, vayanaTypes.AuthRequest{
		HandleType:          "email",
		Handle:              email,
		Password:            password,
		TokenDurationInMins: 360,
	}, &resp, false, true)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf("%s", resp.Error.Message)
	}
	c.tokenLock.Lock()
	c.token = resp.Data.Token
	fmt.Println(resp.Data.Token)
	c.tokenExpiresAt = time.Now().Add(359 * time.Minute)
	c.tokenLock.Unlock()
	return nil
}

func (c *client) Logout() error {
	err, _ := c.sendAuthorizedRequest(http.MethodPost, vayanaTypes.Logout, nil, nil, false, true)
	if err != nil {
		return err
	}
	c.tokenLock.Lock()
	c.token = ""
	c.tokenExpiresAt = time.Time{}
	c.tokenLock.Unlock()
	return nil
}

func (c *client) CreateEWaybill(ewb types.EWBCreateRequest) (*types.EWBCreateResponse, error) {
	endpoint := "/basic/ewb/v1.0/v1.03/gen-ewb"
	resp := &types.EWBCreateResponse{}
	err, vErr := c.sendAuthorizedRequest(http.MethodPost, endpoint, ewb, resp, false, false)
	if err != nil {
		if vErr != nil && vErr.IsEWBError() {
			messages := strings.Join(vErr.GetErrorMessages(), "; ")
			err = fmt.Errorf("%s", messages)
		}
		return resp, fmt.Errorf("failed to create ewaybill: %s", err)
	}
	return resp, nil
}

func (c *client) CancelEWaybill(cancel types.EWBCancelRequest) (*types.EWBCancelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetEWayBill(ewbNo string) (*types.EWBGetResponse, error) {
	endpoint := fmt.Sprintf("/basic/ewb/v1.0/v1.03/%s", ewbNo)
	resp := &types.EWBGetResponse{}
	err, _ := c.sendAuthorizedRequest(http.MethodGet, endpoint, nil, resp, false, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

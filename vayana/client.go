package vayana

import (
	"fmt"
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/vayana/encription"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"github.com/gogotchuri/go-validator"
	"net/http"
	"sync"
	"time"
)

var _ GoGST.Client = &client{}

type client struct {
	validator      *validator.Validate
	validationLock *sync.Mutex

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
		return NewClient(vayanaTypes.BaseProductionAPI, vayanaTypes.TheodoreBaseProductionAPI, organizationID, vayanaTypes.ProductionPublicKey, rek, vayanaTypes.DefaultVersion), nil
	}
	return NewClient(vayanaTypes.BaseSandboxAPI, vayanaTypes.TheodoreBaseSandboxAPI, organizationID, vayanaTypes.SandboxPublicKey, rek, vayanaTypes.DefaultVersion), nil
}

func NewClient(baseURL, theodoreBaseURL, organizationID, publicKey, rek, version string) GoGST.Client {
	return &client{
		validator:       validator.New(),
		validationLock:  &sync.Mutex{},
		apiBaseURL:      baseURL,
		theodoreBaseURL: theodoreBaseURL,
		apiVersion:      version,

		publicKey: publicKey,
		rek:       rek,

		organizationID: organizationID,
		httpClient:     &http.Client{},
	}
}

func (c *client) CreateGSPClient(gstin, username, password string) (GoGST.GSPClient, error) {
	return &gspClient{
		validator:      c.validator,
		validationLock: c.validationLock,
		theodoreClient: c,
		httpClient:     &http.Client{},
		creatorGSTIN:   gstin,
		username:       username,
		password:       password,
	}, nil
}

func (c *client) CreateGSPEInvoicesClient(gstin, username, password string) (GoGST.GSPEInvoiceClient, error) {
	return &gspClient{
		validator:      c.validator,
		validationLock: c.validationLock,

		theodoreClient: c,
		httpClient:     &http.Client{},
		creatorGSTIN:   gstin,
		username:       username,
		password:       password,
	}, nil
}

func (c *client) SetActiveToken(token string) {
	c.tokenLock.Lock()
	c.token = token
	c.tokenExpiresAt = time.Now().Add(60 * time.Minute)
	c.tokenLock.Unlock()
}

func (c *client) Ping() error {
	return c.sendRequest(request{
		method:   http.MethodGet,
		endpoint: vayanaTypes.HealthCheck,
	}, false)
}

func (c *client) Authenticate(email, password string) error {
	resp := vayanaTypes.AuthResponse{}
	err := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: vayanaTypes.AuthTokens,
		body: vayanaTypes.AuthRequest{
			HandleType:          "email",
			Handle:              email,
			Password:            password,
			TokenDurationInMins: 60,
		},
		dest: &resp,
	}, false)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf("%s", resp.Error.Message)
	}
	c.tokenLock.Lock()
	c.token = resp.Data.Token
	fmt.Println(resp.Data.Token)
	c.tokenExpiresAt = time.Now().Add(60 * time.Minute)
	c.tokenLock.Unlock()
	return nil
}

func (c *client) Logout() error {
	err := c.sendRequest(request{
		method:   http.MethodPost,
		endpoint: vayanaTypes.Logout,
	}, true)
	if err != nil {
		return err
	}
	c.tokenLock.Lock()
	c.token = ""
	c.tokenExpiresAt = time.Time{}
	c.tokenLock.Unlock()
	return nil
}

package vayana

import (
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
	"time"
)

func (c *client) isAuthenticated() (bool, error) {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()
	if c.token == "" {
		return false, fmt.Errorf("token is empty")
	}
	if c.tokenExpiresAt.Before(time.Now().Add(1 * time.Minute)) {
		err := c.refreshToken()
		if err != nil {
			c.token = ""
			return false, err
		}
	}
	return true, nil
}

func (c *client) refreshToken() error {
	req := vayanaTypes.AuthResponse{}
	if err, _ := c.sendAuthorizedRequest(http.MethodPut, vayanaTypes.AuthTokens, nil, &req, false, true); err != nil {
		return err
	}
	if req.Error != nil {
		return fmt.Errorf("%s", req.Error.Message)
	}
	c.token = req.Data.Token
	c.tokenExpiresAt = time.Now().Add(359 * time.Minute)
	return nil
}

func (c *client) setTokens(req vayanaTypes.AuthResponse) {
	c.token = req.Data.Token
	c.tokenExpiresAt = time.Unix(req.Data.Expiry, 0)
}

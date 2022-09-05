package vayana

import (
	"fmt"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"net/http"
	"time"
)

func (c *client) IsAuthenticated() (bool, error) {
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

// TODO test
func (c *client) refreshToken() error {
	res := &vayanaTypes.AuthResponse{}
	if err := c.sendRequest(request{
		method:   http.MethodPut,
		endpoint: vayanaTypes.AuthTokens,
		dest:     res,
	}, true); err != nil {
		return err
	}
	if res.Error != nil {
		return fmt.Errorf("%s", res.Error.Message)
	}
	c.token = res.Data.Token
	c.tokenExpiresAt = time.Now().Add(359 * time.Minute)
	return nil
}

func (c *client) setTokens(req vayanaTypes.AuthResponse) {
	c.token = req.Data.Token
	c.tokenExpiresAt = time.Unix(req.Data.Expiry, 0)
}

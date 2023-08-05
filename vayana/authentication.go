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

func (c *client) refreshToken() error {
	res := &vayanaTypes.AuthResponse{}
	req, err := c.newRequest(http.MethodPut, c.getEndpointURL(vayanaTypes.AuthTokens, true), nil, false)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	err, _ = c.send(req, res, false)
	if res.Error != nil {
		return fmt.Errorf("%s", res.Error.Message)
	}
	c.token = res.Data.Token
	c.tokenExpiresAt = time.Now().Add(tokenDurationMin * time.Minute)
	return nil
}

func (c *client) setTokens(req vayanaTypes.AuthResponse) {
	c.token = req.Data.Token
	c.tokenExpiresAt = time.Unix(req.Data.Expiry, 0)
}

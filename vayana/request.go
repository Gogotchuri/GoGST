package vayana

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gogotchuri/GoGST/vayana/encription"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
)

type request struct {
	method   string
	endpoint string
	body     interface{}
	dest     interface{}
}

func (c *client) getEndpointURL(endpoint string, theodore bool) string {
	if endpoint == vayanaTypes.HealthCheck {
		return fmt.Sprintf("%s%s", c.theodoreBaseURL, endpoint)
	} else if theodore {
		return fmt.Sprintf("%s/%s%s", c.theodoreBaseURL, c.apiVersion, endpoint)
	}
	return fmt.Sprintf("%s%s", c.apiBaseURL, endpoint)
}

func (c *client) sendRequest(r request, authenticated bool) error {
	req, err := c.newRequest(r.method, c.getEndpointURL(r.endpoint, true), r.body, false)
	if err != nil {
		return err
	}
	if authenticated {
		if ok, err := c.IsAuthenticated(); !ok {
			return fmt.Errorf("token is empty, athenticate first. %s", err.Error())
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
		req.Header.Set("X-FLYNN-N-USER-TOKEN", c.token)
		req.Header.Set("X-FLYNN-N-ORG-ID", c.organizationID)
	}
	err, _ = c.send(req, r.dest, false)
	return err
}

// send makes a request to the API, the response body will be unmarshalled into v,
// which should be correct struct for the given request body passed by reference or
// it can be an io.Writer, in which case the response bytes will be written to it
func (c *client) send(req *http.Request, dest interface{}, decrypt bool) (error, *vayanaTypes.Error) {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, err = io.ReadAll(resp.Body)
		if err == nil && len(data) > 0 {
			var vayanaError vayanaTypes.ErrorResponse
			err = json.Unmarshal(data, &vayanaError)
			if err == nil {
				if vayanaError.Error.IsTokenExpired() {
					return vayanaTypes.ErrorTokenExpired, &vayanaError.Error
				}
				return fmt.Errorf("%d: %s \n%s", resp.StatusCode, "Request Failed", string(data)), &vayanaError.Error
			}
			return fmt.Errorf("%d: %s \n%s", resp.StatusCode, "Request Failed", string(data)), nil
		}
		return fmt.Errorf("%d: %s", resp.StatusCode, "Request Failed"), nil
	}
	if dest == nil {
		return nil, nil
	}
	if w, ok := dest.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil, nil
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	if decrypt {
		data, err = c.decryptData(data)
		if err != nil {
			return err, nil
		}
	}
	buf := bytes.NewBuffer(data)
	return json.NewDecoder(buf).Decode(dest), nil
}

// newRequest constructs a new http.Request, Marshal payload to json bytes
func (c *client) newRequest(method, url string, payload interface{}, encrypt bool) (*http.Request, error) {
	var buf io.Reader
	var encryptedRek, encryptedData []byte
	if payload != nil {
		b, err := json.Marshal(&payload)
		fmt.Println(string(b))
		if err != nil {
			return nil, err
		}
		if encrypt {
			encryptedRek, encryptedData, err = c.encryptRekAndData(b)
			if err != nil {
				return nil, err
			}
			buf = bytes.NewBuffer(encryptedData)
		} else {
			buf = bytes.NewBuffer(b)
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	if encrypt {
		req.Header.Set("X-FLYNN-S-REK", string(encryptedRek))
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	return req, nil
}

func (c *client) decryptData(data []byte) ([]byte, error) {
	return encription.AESECBEncryption(c.rek, string(data))
}

func (c *client) encryptRekAndData(data []byte) (encryptedRek []byte, encryptedData []byte, err error) {
	eRek, err := encription.AESECBEncryption(c.publicKey, c.rek)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encrypt rek: %s", err.Error())
	}
	eData, err := encription.AESECBEncryption(c.rek, string(data))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encrypt data: %s", err.Error())
	}
	//Encode the encrypted data and rek in base64
	base64.StdEncoding.Encode(encryptedRek, eRek)
	base64.StdEncoding.Encode(encryptedData, eData)
	return encryptedRek, encryptedData, nil
}

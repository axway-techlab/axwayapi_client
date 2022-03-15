package axwayapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	if host == nil {
		return nil, fmt.Errorf("a host must be given")
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL: *host,
	}

	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, expect... int) ([]byte, error) {
	username := c.Auth.Username
	password := c.Auth.Password

	if len(expect) == 0 {
		// when no expected status code is given, 200 is implied.
		expect = append(expect, http.StatusOK)
	}

	req.SetBasicAuth(username,password)
	if _, ok := req.Header["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	for _, expected := range expect {
		if res.StatusCode == expected {
			return body, err
		}
	}
	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
}

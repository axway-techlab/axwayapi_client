package axwayapi

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
func NewClient(host, username, password string, proxy *url.URL, insecureSkipVerify bool) (*Client, error) {
	if host == "" {
		return nil, fmt.Errorf("a host must be given")
	}

	c := Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
			},
		},
		HostURL: host,
	}

	if username == "" || password == "" {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: username,
		Password: password,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, expect ...int) ([]byte, error) {
	username := c.Auth.Username
	password := c.Auth.Password

	if len(expect) == 0 {
		// when no expected status code is given, some classical 20x are implied.
		expect = []int{http.StatusOK, http.StatusNoContent, http.StatusCreated}
	}

	req.SetBasicAuth(username, password)
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

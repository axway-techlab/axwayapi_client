package axwayapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetConfig - Returns a specifc Config
func (c *Client) GetConfig() (*Config, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/config", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	Config := Config{}
	err = json.Unmarshal(body, &Config)
	if err != nil {
		return nil, err
	}

	return &Config, nil
}

// UpdateConfig - Updates an Config
func (c *Client) UpdateConfig(config *Config) (*Config, error) {
	rb, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/config", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	Config := Config{}
	err = json.Unmarshal(body, &Config)
	if err != nil {
		return nil, err
	}

	return &Config, nil
}


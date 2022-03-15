package axwayapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

func (c *Client) CreateOrg(org *Org) error {
	rb, err := json.Marshal(org)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/organizations", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, 201)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &org)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetOrg(id string) (*Org, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	Org := Org{}
	err = json.Unmarshal(body, &Org)
	if err != nil {
		return nil, err
	}

	return &Org, nil
}

func (c *Client) ListOrgs() (ret []Org, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	orgs := make([]Org, 0)
	err = json.Unmarshal(body, &orgs)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (c *Client) UpdateOrg(org *Org) error {
	rb, err := json.Marshal(org)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s", c.HostURL, org.Id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &org)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateOrgImage(id string, image string) error {
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		return err
	}
	img, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return err
	}
	fw.Write(img)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/image/", c.HostURL, id), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOrg(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, 204)
	if err != nil {
		return err
	}

	return nil
}

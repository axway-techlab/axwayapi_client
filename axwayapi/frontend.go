package axwayapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

func (c *Client) CreateFrontend(frontend *Frontend) (error) {
	rb, err := json.Marshal(frontend)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/proxies", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, 201)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &frontend)
	if err != nil {
		return err
	}

	return nil}

func (c *Client) GetFrontend(id string) (*Frontend, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/proxies/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	frontend := Frontend{}
	err = json.Unmarshal(body, &frontend)
	if err != nil {
		return nil, err
	}

	return &frontend, nil
}

func (c *Client) ListFrontends() (ret []Frontend, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/proxies/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	users := make([]Frontend, 0)
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UpdateFrontend(frontend *Frontend) error {
	rb, err := json.Marshal(frontend)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/proxies/%s", c.HostURL, frontend.Id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &frontend)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PublishFrontend(frontend *Frontend) error {
	return c.operateOnFrontend(frontend, "publish")
}

func (c *Client) UnpublishFrontend(frontend *Frontend) error {
	return c.operateOnFrontend(frontend, "unpublish")
}

func (c *Client) DeprecateFrontend(frontend *Frontend) error {
	return c.operateOnFrontend(frontend, "deprecate")
}

func (c *Client) UndeprecateFrontend(frontend *Frontend) error {
	return c.operateOnFrontend(frontend, "undeprecate")
}

func (c *Client) operateOnFrontend(frontend *Frontend, operation string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/proxies/%s/%s", c.HostURL, frontend.Id, operation), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &frontend)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateFrontendImage(id string, image string) error {
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	img, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return err
	}

	//---
	partHeader := textproto.MIMEHeader{}
	partHeader.Add("Content-Disposition", "form-data; name=\"file\"; filename=\"image.jpg\"")
	partHeader.Add("Content-Type", "image/jpeg")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return err
	}
	_, err = part.Write(img)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	//---
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/proxies/%s/image/", c.HostURL, id), body)
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

func (c *Client) DeleteFrontend(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/proxies/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

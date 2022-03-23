package axwayapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

//---
type Backend struct {
	Id                    string                 `json:"id,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Summary               string                 `json:"summary,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Version               string                 `json:"version,omitempty"`
	BasePath              string                 `json:"basePath,omitempty"`
	ResourcePath          string                 `json:"resourcePath,omitempty"`
	Consumes              []string               `json:"consumes,omitempty"`
	Produces              []string               `json:"produces,omitempty"`
	Integral              bool                   `json:"integral,omitempty"`
	CreatedOn             int                    `json:"createdOn,omitempty"`
	CreatedBy             string                 `json:"createdBy,omitempty"`
	OrganizationId        string                 `json:"organizationId,omitempty"`
	ServiceType           string                 `json:"serviceType,omitempty"`
	HasOriginalDefinition bool                   `json:"hasOriginalDefinition,omitempty"`
	ImportUrl             string                 `json:"importUrl,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	Models                map[string]interface{} `json:"models,omitempty"`
}

func (c *Client) CreateBackend(orgid, name, apitype, file string) (*Backend, error) {
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//---
	partHeader := textproto.MIMEHeader{}
	partHeader.Add("Content-Disposition", "form-data; name=\"file\"; filename=\"image.jpg\"")
	partHeader.Add("Content-Type", "application/json")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return nil,err
	}
	part.Write([]byte(file))

	//---
	err = writer.WriteField("name", name)
	if err != nil {
		return nil,err
	}
	//---
	err = writer.WriteField("type", apitype)
	if err != nil {
		return nil,err
	}

	//--
	err = writer.WriteField("organizationId", orgid)
	if err != nil {
		return nil,err
	}

	writer.Close()

	//---
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apirepo/import", c.HostURL), body)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rbody, err := c.doRequest(req)
	if err != nil {
		return nil,err
	}

	backend := Backend{}
	err = json.Unmarshal(rbody, &backend)
	if err != nil {
		return nil, err
	}
	return &backend, nil
}

func (c *Client) GetBackend(id string) (*Backend, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apirepo/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	backend := Backend{}
	err = json.Unmarshal(body, &backend)
	if err != nil {
		return nil, err
	}

	return &backend, nil
}

func (c *Client) ListBackends() (ret []Backend, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apirepo/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	users := make([]Backend, 0)
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UpdateBackend(backend *Backend) error {
	rb, err := json.Marshal(backend)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apirepo/%s", c.HostURL, backend.Id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &backend)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteBackend(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apirepo/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

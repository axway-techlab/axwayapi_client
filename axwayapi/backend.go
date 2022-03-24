package axwayapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) CreateBackend(orgid, name, apitype, file string) (ret *Backend, err error) {
	mpart :=
		newMultiPart().
			addFile(filePart{
				data:        []byte(file),
				contentType: "application/json",
				fileName:    "swagger.json",
			}).
			addField("name", name).
			addField("type", apitype).
			addField("organizationId", orgid)

	body, err := c.sendParts("apirepo/import", mpart)
	if err != nil {
		return nil, err
	}

	ret = &Backend{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) GetBackend(id string) (ret *Backend, err error) {
	ret = &Backend{}
	err = c.get(ret, fmt.Sprintf("apirepo/%s", id))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) ListBackends() (ret []Backend, err error) {
	err = c.get(ret, "apirepo")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) UpdateBackend(backend *Backend) error {
	return c.put(backend, fmt.Sprintf("apirepo/%s", backend.Id))
}

func (c *Client) DeleteBackend(id string) error {
	return c.delete(fmt.Sprintf("apirepo/%s", id))
}

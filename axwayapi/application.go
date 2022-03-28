package axwayapi

import (
	"fmt"
	"net/http"
	"sort"
)

type Application struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	OrganizationId string `json:"organizationId"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	State          string `json:"state,omitempty"`
	//	Tags           map[string]string `json:"tags,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	CreatedOn int       `json:"createdOn,omitempty"`
	ManagedBy []string  `json:"managedBy,omitempty"`
	Apis      *[]string `json:"apis,omitempty"`
}
func (this *Application) GetId() string {
	return this.Id
} 

type ApiLink struct {
	Id        string `json:"id,omitempty"`
	ApiId     string `json:"apiId"`
	Enabled   bool   `json:"enabled,omitempty"`
	State     string `json:"state,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedOn int    `json:"createdOn,omitempty"`
}

type ApiKey struct {
	Id            string   `json:"id,omitempty"`
	ApplicationId string   `json:"applicationId"`
	Enabled       bool     `json:"enabled,omitempty"`
	Secret        string   `json:"secret,omitempty"`
	CreatedBy     string   `json:"createdBy,omitempty"`
	CreatedOn     int      `json:"createdOn,omitempty"`
	DeletedOn     int      `json:"deletedOn,omitempty"`
	CorsOrigins   []string `json:"corsOrigins"`
}

func (c *Client) CreateApplication(application *Application) error {
	// At creation time, this field must be set to the empty array...
	application.Apis = &[]string{}
	return c.post(application, "applications")
}

func (c *Client) GetApplication(id string) (ret *Application, err error) {
	ret = &Application{}
	err = c.get(ret, fmt.Sprintf("applications/%s", id))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) ListApplications() (ret []Application, err error) {
	err = c.get(ret, "applications")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) DeleteApplication(id string) error {
	return c.delete(fmt.Sprintf("applications/%s", id))
}

func (c *Client) UpdateApplication(application *Application) (err error) {
	// At update time, this field must be omitted...
	application.Apis = nil
	return c.put(application, fmt.Sprintf("applications/%s", application.Id))
}

func (c *Client) UpdateApplicationImage(id string, image string) error {
	return c.updateImage(fmt.Sprintf("applications/%s/image/", id), image)
}

// ApiKeys <-> Applications
func (c *Client) AddApiKeyToApplication(application *Application, apiKey *ApiKey) error {
	return c.post(
		apiKey,
		fmt.Sprintf("applications/%s/apikeys", application.Id),
		http.StatusCreated,
	)
}

func (c *Client) DeleteApiKeyFromApplication(appId string, apiKeyId string) error {
	return c.delete(
		fmt.Sprintf("applications/%s/apikeys/%s", appId, apiKeyId),
		http.StatusNoContent)
}

func (c *Client) ListApiKeysInApplication(appId string) ([]string, error) {
	var apis []ApiKey
	c.get(
		&apis,
		fmt.Sprintf("applications/%s/apikeys", appId),
		http.StatusOK, http.StatusBadRequest,
	)
	ids := make([]string, len(apis))
	for i, a := range apis {
		ids[i] = a.Id
	}
	sort.Strings(ids)
	return ids, nil
}

// Apis <-> Applications
func (c *Client) AddApiToApplication(application *Application, apiId string) error {
	link := &ApiLink{ApiId: apiId, Enabled: true}

	return c.post(link,
		fmt.Sprintf("applications/%s/apis", application.Id),
		http.StatusCreated,
	)
}

func (c *Client) DeleteApiFromApplication(application *Application, apiId string) error {
	return c.delete(fmt.Sprintf("applications/%s/apis/%s", application.Id, apiId))
}

func (c *Client) ListApisInApplication(appId string) ([]string, error) {
	var apis []ApiLink
	err := c.get(
		&apis,
		fmt.Sprintf("applications/%s/apis", appId),
		http.StatusOK, http.StatusBadRequest,
	)
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(apis))
	for i, a := range apis {
		ids[i] = a.ApiId
	}
	sort.Strings(ids)
	return ids, nil
}

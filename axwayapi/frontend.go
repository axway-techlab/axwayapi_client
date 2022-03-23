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

type Frontend struct {
	Id                     string                     `json:"id,omitempty"`
	OrganizationId         string                     `json:"organizationId"`
	ApiId                  string                     `json:"apiId"`
	Name                   string                     `json:"name"`
	Version                string                     `json:"version"`
	ApiRoutingKey          string                     `json:"apiRoutingKey"`
	Vhost                  string                     `json:"vhost"`
	Path                   string                     `json:"path,omitempty"`
	DescriptionType        string                     `json:"descriptionType,omitempty"`
	DescriptionManual      string                     `json:"descriptionManual,omitempty"`
	DescriptionMarkdown    string                     `json:"descriptionMarkdown,omitempty"`
	DescriptionUrl         string                     `json:"descriptionUrl,omitempty"`
	Summary                string                     `json:"summary,omitempty"`
	Retired                bool                       `json:"retired,omitempty"`
	Expired                bool                       `json:"expired,omitempty"`
	Image                  string                     `json:"image,omitempty"`
	RetirementDate         int                        `json:"retirementDate,omitempty"`
	Deprecated             bool                       `json:"deprecated,omitempty"`
	State                  string                     `json:"state,omitempty"`
	CorsProfiles           []CorsProfile              `json:"corsProfiles,omitempty"`
	SecurityProfiles       []SecurityProfile          `json:"securityProfiles,omitempty"`
	AuthenticationProfiles []AuthenticationProfile    `json:"authenticationProfiles,omitempty"`
	InboundProfiles        map[string]InboundProfile  `json:"inboundProfiles,omitempty"`
	OutboundProfiles       map[string]OutboundProfile `json:"outboundProfiles,omitempty"`
	ServiceProfiles        map[string]ServiceProfile  `json:"serviceProfiles,omitempty"`
	CACerts                []CACert                   `json:"caCerts,omitempty"`
	Tags                   map[string][]string        `json:"tags,omitempty"`
	CustomProperties       map[string]interface{}     `json:"customProperties,omitempty"`
	CreatedBy              string                     `json:"createdBy,omitempty"`
	CreatedOn              int                        `json:"createdOn,omitempty"`
}

type CorsProfile struct {
	//ValueObject
	Name               string   `json:"name"`
	IsDefault          bool     `json:"isDefault"`
	Origins            []string `json:"origins"`
	AllowedHeaders     []string `json:"allowedHeaders"`
	ExposedHeaders     []string `json:"exposedHeaders"`
	SupportCredentials bool     `json:"supportCredentials"`
	MaxAgeSeconds      int      `json:"maxAgeSeconds,omitempty"`
}
type SecurityProfile struct {
	//ValueObject
	Name      string   `json:"name"`
	IsDefault bool     `json:"isDefault"`
	Devices   []Device `json:"devices"`
}
type Device struct {
	//ValueObject
	Name       string                 `json:"name,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Order      int                    `json:"order,omitempty"`
	Properties map[string]interface{} `json:"properties"`
}
type AuthenticationProfile struct {
	//ValueObject
	Name       string                 `json:"name,omitempty"`
	Type       string                 `json:"type,omitempty"`
	IsDefault  bool                   `json:"isDefault,omitempty"`
	Parameters map[string]interface{} `json:"parameters"`
}
type InboundProfile struct {
	//ValueObject
	SecurityProfile string `json:"securityProfile,omitempty"`
	CorsProfile     string `json:"corsProfile,omitempty"`
	MonitorAPI      bool   `json:"monitorAPI,omitempty"`
	MonitorSubject  string `json:"monitorSubject,omitempty"`
}
type OutboundProfile struct {
	//ValueObject
	AuthenticationProfile string       `json:"authenticationProfile,omitempty"`
	RouteType             string       `json:"routeType,omitempty"`
	RequestPolicy         string       `json:"requestPolicy,omitempty"`
	ResponsePolicy        string       `json:"responsePolicy,omitempty"`
	RoutePolicy           string       `json:"routePolicy,omitempty"`
	FaultHandlerPolicy    string       `json:"faultHandlerPolicy,omitempty"`
	ApiId                 string       `json:"apiId,omitempty"`
	ApiMethodId           string       `json:"apiMethodId,omitempty"`
	Parameters            []ParamValue `json:"parameters"`
}
type ParamValue struct {
	//ValueObject
	Name       string `json:"name,omitempty"`
	ParamType  string `json:"paramType,omitempty"`
	Type       string `json:"type,omitempty"`
	Format     string `json:"format,omitempty"`
	Value      string `json:"value,omitempty"`
	Required   bool   `json:"required,omitempty"`
	Exclude    bool   `json:"exclude,omitempty"`
	Additional bool   `json:"additional,omitempty"`
}
type ServiceProfile struct {
	//ValueObject
	ApiId    string `json:"apiId"`
	BasePath string `json:"basePath"`
}
type CACert struct {
	//ValueObject
	CertBlob           string `json:"certBlob"`
	Name               string `json:"name"`
	Alias              string `json:"alias"`
	Subject            string `json:"subject"`
	Issuer             string `json:"issuer"`
	Version            int    `json:"version"`
	NotValidBefore     int    `json:"notValidBefore"`
	NotValidAfter      int    `json:"notValidAfter"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	Sha1Fingerprint    string `json:"sha1Fingerprint"`
	Md5Fingerprint     string `json:"md5Fingerprint"`
	Expired            bool   `json:"expired"`
	NotYetValid        bool   `json:"notYetValid"`
	Inbound            bool   `json:"inbound"`
	Outbound           bool   `json:"outbound"`
}

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

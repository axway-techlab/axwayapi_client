package axwayapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

type User struct {
	Id             string                 `json:"id,omitempty"`
	OrganizationId string                 `json:"organizationId"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description,omitempty"`
	LoginName      string                 `json:"loginName"`
	Email          string                 `json:"email,omitempty"`
	Phone          string                 `json:"phone,omitempty"`
	Mobile         string                 `json:"mobile,omitempty"`
	Role           string                 `json:"role"`
	Image          string                 `json:"image,omitempty"`
	Enabled        bool                   `json:"enabled"`
	CreatedOn      int                    `json:"createdOn,omitempty"`
	State          string                 `json:"state,omitempty"`
	Type           string                 `json:"type,omitempty"`
	Orgs2Role      map[string]string      `json:"orgs2Role,omitempty"`
	AuthAttrs      map[string]interface{} `json:"authAttrs,omitempty"`
	Dn             string                 `json:"dn,omitempty"`
}

func (c *Client) CreateUser(user *User) error {
	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, 201)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUser(id string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) ListUsers() (ret []User, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UpdateUser(user *User) error {
	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s", c.HostURL, user.Id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateUserImage(id string, image string) error {
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
	part.Write(img)
	writer.Close()

	//---
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/%s/image/", c.HostURL, id), body)
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

func (c *Client) SetPassword(userId string, pwd string) error {
	formData := url.Values{"newPassword": {pwd}}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/%s/changepassword", c.HostURL, userId), strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := c.doRequest(req, 204)
	if err != nil {
		return fmt.Errorf("got body '%s', and error '%v'", string(body), err)
	}
	return nil
}

func (c *Client) DeleteUser(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, 204)
	if err != nil {
		return err
	}

	return nil
}

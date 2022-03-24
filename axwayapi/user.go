package axwayapi

import (
	"fmt"
	"net/http"
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
	return c.post(user, "users")
}

func (c *Client) GetUser(id string) (ret *User, err error) {
	ret = &User{}
	err = c.get(ret, fmt.Sprintf("users/%s", id))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) ListUsers() (ret []User, err error) {
	err = c.get(ret, "users")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) DeleteUser(id string) error {
	return c.delete(fmt.Sprintf("users/%s", id))
}

func (c *Client) UpdateUser(user *User) error {
	return c.put(user, fmt.Sprintf("users/%s", user.Id))
}

func (c *Client) UpdateUserImage(id string, image string) error {
	return c.updateImage(fmt.Sprintf("users/%s/image", id), image)
}

func (c *Client) SetPassword(userId string, pwd string) error {
	formData := url.Values{"newPassword": {pwd}}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/%s/changepassword", c.HostURL, userId), strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := c.doRequest(req, http.StatusNoContent)
	if err != nil {
		return fmt.Errorf("got body '%s', and error '%v'", string(body), err)
	}
	return nil
}

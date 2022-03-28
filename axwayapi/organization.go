package axwayapi

import (
	"fmt"
)

type Org struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Email          string `json:"email"`
	Restricted     bool   `json:"restricted,omitempty"`
	VirtualHost    string `json:"virtualHost"`
	Phone          string `json:"phone"`
	Enabled        bool   `json:"enabled"`
	Development    bool   `json:"development"`
	Dn             string `json:"dn,omitempty"`
	CreatedOn      int    `json:"createdOn,omitempty"`
	StartTrialDate int    `json:"startTrialDate,omitempty"`
	EndTrialDate   int    `json:"endTrialDate,omitempty"`
	TrialDuration  int    `json:"trialDuration,omitempty"`
	IsTrial        bool   `json:"isTrial,omitempty"`
}
func (this *Org) GetId() string {
	return this.Id
} 

func (c *Client) CreateOrg(org *Org) error {
	return c.post(org, "organizations")
}

func (c *Client) GetOrg(id string) (ret *Org, err error) {
	ret = &Org{}
	err = c.get(ret, fmt.Sprintf("organizations/%s", id))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) ListOrgs() (ret []Org, err error) {
	err = c.get(ret, "organizations")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) DeleteOrg(id string) error {
	return c.delete(fmt.Sprintf("organizations/%s", id))
}

func (c *Client) UpdateOrg(org *Org) error {
	return c.put(org, fmt.Sprintf("organizations/%s", org.Id))
}

func (c *Client) UpdateOrgImage(id string, image string) error {
	return c.updateImage(fmt.Sprintf("organizations/%s/image/", id), image)
}

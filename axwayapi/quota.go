package axwayapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Quota struct {
	Id           string       `json:"id,omitempty"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Description  string       `json:"description,omitempty"`
	System       bool         `json:"system"`
	Restrictions []Constraint `json:"restrictions,omitempty"`
}

type Constraint struct {
	bareConstraint
	Config interface{} `json:"config"`
}
type bareConstraint struct {
	Api    string `json:"api"`
	Method string `json:"method"`
	Type   string `json:"type"`
}

func (cc *Constraint) UnmarshalJSON(b []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &cc.bareConstraint)
	if err != nil {
		return err
	}
	m := make(map[string]string)
	err = json.Unmarshal(raw["config"], &m)
	fmt.Printf("Map : %#+v\n", m)
	if err != nil {
		return err
	}
	per, err := strconv.Atoi(m["per"])
	if err != nil {
		return err
	}
	switch cc.Type {
	case "throttlemb":
		n, err := strconv.Atoi(m["mb"])
		if err != nil {
			return err
		}
		cc.Config = ConstraintConfigMb{
			Mb:     n,
			Per:    per,
			Period: m["period"],
		}
	case "throttle":
		n, err := strconv.Atoi(m["msg"])
		if err != nil {
			return err
		}
		cc.Config = ConstraintConfigMsg{
			Msg:    n,
			Per:    per,
			Period: m["period"],
		}
	default:
		return fmt.Errorf("cannot unmarshall constraint of type %s", cc.Type)
	}
	return err
}

type ConstraintConfigMb struct {
	Per    int    `json:"per"`
	Period string `json:"period"`
	Mb     int    `json:"mb"`
}
type ConstraintConfigMsg struct {
	Per    int    `json:"per"`
	Period string `json:"period"`
	Msg    int    `json:"msg"`
}

func (c *Client) CreateQuota(quota *Quota) error {
	return c.post(quota, "quotas")
}

func (c *Client) GetQuota(id string) (ret *Quota, err error) {
	ret = &Quota{}
	err = c.get(ret, fmt.Sprintf("quotas/%s", id))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) ListQuotas() (ret []Quota, err error) {
	err = c.get(ret, "quotas")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) DeleteQuota(id string) error {
	return c.delete(fmt.Sprintf("quotas/%s", id))
}

func (c *Client) UpdateQuota(quota *Quota) (err error) {
	return c.put(quota, fmt.Sprintf("quotas/%s", quota.Id))
}

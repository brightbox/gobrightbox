package brightbox

import (
	"time"
)

type ServerGroup struct {
	Resource
	Name           string
	CreatedAt      *time.Time `json:"created_at"`
	Description    string
	Default        bool
	Account        Account `json:"account"`
	Servers        []Server
	FirewallPolicy FirewallPolicy
}

type ServerGroupOptions struct {
	Identifier  string  `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (c *Client) ServerGroups() (*[]ServerGroup, error) {
	groups := new([]ServerGroup)
	_, err := c.MakeApiRequest("GET", "/1.0/server_groups", nil, groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}

func (c *Client) ServerGroup(identifier string) (*ServerGroup, error) {
	group := new(ServerGroup)
	_, err := c.MakeApiRequest("GET", "/1.0/server_groups/"+identifier, nil, group)
	if err != nil {
		return nil, err
	}
	return group, err
}

func (c *Client) CreateServerGroup(newServerGroup *ServerGroupOptions) (*ServerGroup, error) {
	group := new(ServerGroup)
	_, err := c.MakeApiRequest("POST", "/1.0/server_groups", newServerGroup, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (c *Client) UpdateServerGroup(updateServerGroup *ServerGroupOptions) (*ServerGroup, error) {
	group := new(ServerGroup)
	_, err := c.MakeApiRequest("PUT", "/1.0/server_groups/"+updateServerGroup.Identifier, updateServerGroup, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (c *Client) DestroyServerGroup(identifier string) error {
	_, err := c.MakeApiRequest("DELETE", "/1.0/server_groups/"+identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

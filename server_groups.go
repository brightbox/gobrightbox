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

type serverGroupMemberOptions struct {
	Servers     []serverGroupMember `json:"servers"`
	Destination string              `json:"destination,omitempty"`
}
type serverGroupMember struct {
	Server string `json:"server,omitempty"`
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

func (c *Client) AddServersToServerGroup(identifier string, serverIds []string) (*ServerGroup, error) {
	group := new(ServerGroup)
	opts := new(serverGroupMemberOptions)
	for _, id := range serverIds {
		opts.Servers = append(opts.Servers, serverGroupMember{Server: id})
	}
	_, err := c.MakeApiRequest("POST", "/1.0/server_groups/"+identifier+"/add_servers", opts, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (c *Client) RemoveServersFromServerGroup(identifier string, serverIds []string) (*ServerGroup, error) {
	group := new(ServerGroup)
	opts := new(serverGroupMemberOptions)
	for _, id := range serverIds {
		opts.Servers = append(opts.Servers, serverGroupMember{Server: id})
	}
	_, err := c.MakeApiRequest("POST", "/1.0/server_groups/"+identifier+"/remove_servers", opts, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (c *Client) MoveServersToServerGroup(src string, dst string, serverIds []string) (*ServerGroup, error) {
	group := new(ServerGroup)
	opts := serverGroupMemberOptions{Destination: dst}
	for _, id := range serverIds {
		opts.Servers = append(opts.Servers, serverGroupMember{Server: id})
	}
	_, err := c.MakeApiRequest("POST", "/1.0/server_groups/"+src+"/move_servers", opts, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

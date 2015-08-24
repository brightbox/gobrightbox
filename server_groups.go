package brightbox

import (
	"time"
)

type ServerGroup struct {
	Resource
	Name        string
	CreatedAt   *time.Time `json:"created_at"`
	Description string
	Default     bool
	Account     *Account `json:"account"`
	Servers     []*Server
}

func (c *Client) ServerGroups() (*[]ServerGroup, error) {
	groups := new([]ServerGroup)
	_, err := c.MakeApiRequest("GET", "/1.0/server_groups", nil, groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}

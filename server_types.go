package gobrightbox

import (
	"fmt"
)

// ServerType represents a Server Type
// https://api.gb1.brightbox.com/1.0/#server_type
type ServerType struct {
	ID       string
	Name     string
	Status   string
	Handle   string
	Cores    int
	RAM      int
	DiskSize int `json:"disk_size"`
}

// ServerTypes retrieves a list of all Server Types
func (c *Client) ServerTypes() ([]ServerType, error) {
	var servertypes []ServerType
	_, err := c.MakeAPIRequest("GET", "/1.0/server_types", nil, &servertypes)
	if err != nil {
		return nil, err
	}
	return servertypes, err
}

// ServerType retrieves a detailed view of one Server Type using an identifier
func (c *Client) ServerType(identifier string) (*ServerType, error) {
	servertype := new(ServerType)
	_, err := c.MakeAPIRequest("GET", "/1.0/server_types/"+identifier, nil, servertype)
	if err != nil {
		return nil, err
	}
	return servertype, err
}

// ServerTypeByHandle retrieves a detailed view of one Server Type using a handle
func (c *Client) ServerTypeByHandle(handle string) (*ServerType, error) {
	servertypes, err := c.ServerTypes()
	if err != nil {
		return nil, err
	}
	for _, servertype := range servertypes {
		if servertype.Handle == handle {
			return &servertype, nil
		}
	}
	return nil, fmt.Errorf("ServerType with handle '%s' doesn't exist", handle)
}

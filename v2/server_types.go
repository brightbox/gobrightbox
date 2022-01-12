package brightbox

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

func (_c *ServerType) APIPath() string {
	return "server_types"
}

func (c *ServerType) HandleString() string {
	return c.Handle
}

type handleable interface {
	queriable
	HandleString() string
}

// ByHandle retrieves a detailed view of a Resource using a handle
func ByHandle[T handleable](q *Client, handle string) (T, error) {
	servertypes, err := All[T](q)
	if err != nil {
		return *new(T), err
	}
	for _, servertype := range servertypes {
		if servertype.HandleString() == handle {
			return servertype, nil
		}
	}
	return *new(T), fmt.Errorf("Resource with handle '%s' doesn't exist", handle)
}

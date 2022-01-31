package brightbox

import (
	"fmt"
)

type handleable interface {
	queriable
	HandleString() string
}

// ByHandle retrieves a detailed view of a Resource using a handle
func ByHandle[T handleable](q *Client, handle string) (*T, error) {
	servertypes, err := All[T](q)
	if err != nil {
		return nil, err
	}
	for _, servertype := range servertypes {
		if servertype.HandleString() == handle {
			return &servertype, nil
		}
	}
	return nil, fmt.Errorf("Resource with handle '%s' doesn't exist", handle)
}

// Code generated by go generate; DO NOT EDIT.

package brightbox

import "context"
import "path"

import "fmt"

const (
	// servertypeAPIPath returns the relative URL path to the ServerType endpoint
	servertypeAPIPath = "server_types"
)

// ServerTypes returns the collection view for ServerType
func (c *Client) ServerTypes(ctx context.Context) ([]ServerType, error) {
	return apiGetCollection[[]ServerType](ctx, c, servertypeAPIPath)
}

// ServerType retrieves a detailed view of one resource
func (c *Client) ServerType(ctx context.Context, identifier string) (*ServerType, error) {
	return apiGet[ServerType](ctx, c, path.Join(servertypeAPIPath, identifier))
}

// ServerType retrieves a detailed view of one resource using a handle
func (c *Client) ServerTypeByHandle(ctx context.Context, handle string) (*ServerType, error) {
	collection, err := c.ServerTypes(ctx)
	if err != nil {
		return nil, err
	}
	for _, instance := range collection {
		if instance.Handle == handle {
			return &instance, nil
		}
	}
	return nil, fmt.Errorf("Resource with handle '%s' doesn't exist", handle)
}

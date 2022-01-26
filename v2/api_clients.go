package brightbox

import (
	"time"
)

// APIClient represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type APIClient struct {
	ID               string
	Name             string
	Description      string
	Secret           string
	PermissionsGroup string     `json:"permissions_group"`
	RevokedAt        *time.Time `json:"revoked_at"`
//	Account          Account
}

// APIClientOptions is used to create and update api clients
type APIClientOptions struct {
	ID               string  `json:"-"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	PermissionsGroup *string `json:"permissions_group,omitempty"`
}

// APIPath returns the relative URL path to the config map collection
func (c APIClient) APIPath() string {
	return "api_clients"
}

// Extract copies a APIClient object to a APIClientOptions object
func (c APIClient) Extract() *APIClientOptions {
	return &APIClientOptions{
		ID:               c.ID,
		Name:             &c.Name,
		Description:      &c.Description,
		PermissionsGroup: &c.PermissionsGroup,
	}
}

// FetchID returns the ID field from a APIClientOptions object
// ID will be blank for create, and set for update
func (c APIClientOptions) FetchID() string {
	return c.ID
}

// LockID returns the path to a lockable APIClient object
func (c APIClient) LockID() string {
	return c.APIPath() + "/" + c.ID
}

package brightbox

import (
	"time"
)

// APIClient represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type APIClient struct {
	ResourceRef
	ID               string
	Name             string
	Description      string
	Secret           string
	PermissionsGroup string     `json:"permissions_group"`
	RevokedAt        *time.Time `json:"revoked_at"`
	Account          *Account
}

// APIClientOptions is used to create and update api clients
type APIClientOptions struct {
	ID               string  `json:"-"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	PermissionsGroup *string `json:"permissions_group,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c APIClient) APIPath() string {
	return "api_clients"
}

// FetchID returns the ID field from the object
func (c APIClient) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c APIClient) PostPath(from *APIClientOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c APIClient) PutPath(from *APIClientOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c APIClient) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c APIClientOptions) OptionID() string {
	return c.ID
}

// ResetPasswordPath returns the relative URL path to reset the password
func (c APIClient) ResetPasswordPath() string {
	return c.APIPath() + "/" + c.FetchID() + "/reset_secret"
}

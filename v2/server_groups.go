package brightbox

import (
	"time"
)

// ServerGroup represents a server group
// https://api.gb1.brightbox.com/1.0/#server_group
type ServerGroup struct {
	ID             string
	Name           string
	CreatedAt      *time.Time `json:"created_at"`
	Description    string
	Default        bool
	Fqdn           string
	Account        *Account
	FirewallPolicy *FirewallPolicy `json:"firewall_policy"`
	Servers        []Server
}

// ServerGroupOptions is used in combination with CreateServerGroup and
// UpdateServerGroup to create and update server groups
type ServerGroupOptions struct {
	ID          string  `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c ServerGroup) APIPath() string {
	return "server_groups"
}

// PostPath returns the relative URL path to POST an object
func (c ServerGroup) PostPath(from *ServerGroupOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c ServerGroup) PutPath(from *ServerGroupOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c ServerGroup) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// FetchID returns the ID field from the object
func (c ServerGroup) FetchID() string {
	return c.ID
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c ServerGroupOptions) OptionID() string {
	return c.ID
}
package brightbox

import (
	"time"
)

// FirewallPolicy represents a firewall policy.
// https://api.gb1.brightbox.com/1.0/#firewall_policy
type FirewallPolicy struct {
	ResourceRef
	ID          string
	Name        string
	Default     bool
	CreatedAt   *time.Time `json:"created_at"`
	Description string
	Account     *Account
	ServerGroup *ServerGroup   `json:"server_group"`
	Rules       []FirewallRule `json:"rules"`
}

// FirewallPolicyOptions is used in conjunction with CreateFirewallPolicy and
// UpdateFirewallPolicy to create and update firewall policies.
type FirewallPolicyOptions struct {
	ID          string  `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ServerGroup *string `json:"server_group,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c FirewallPolicy) APIPath() string {
	return "firewall_policies"
}

// FetchID returns the ID field from the object
func (c FirewallPolicy) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c FirewallPolicy) PostPath(from *FirewallPolicyOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c FirewallPolicy) PutPath(from *FirewallPolicyOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c FirewallPolicy) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c FirewallPolicyOptions) OptionID() string {
	return c.ID
}

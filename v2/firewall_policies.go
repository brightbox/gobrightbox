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

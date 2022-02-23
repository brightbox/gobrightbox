package brightbox

import (
	"time"
)

// ServerGroup represents a server group
// https://api.gb1.brightbox.com/1.0/#server_group
type ServerGroup struct {
	ResourceRef
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

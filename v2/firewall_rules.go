package brightbox

import (
	"time"
)

// FirewallRule represents a firewall rule.
// https://api.gb1.brightbox.com/1.0/#firewall_rule
type FirewallRule struct {
	ID              string
	Source          string          `json:"source"`
	SourcePort      string          `json:"source_port"`
	Destination     string          `json:"destination"`
	DestinationPort string          `json:"destination_port"`
	Protocol        string          `json:"protocol"`
	IcmpTypeName    string          `json:"icmp_type_name"`
	CreatedAt       time.Time       `json:"created_at"`
	Description     string          `json:"description"`
	FirewallPolicy  *FirewallPolicy `json:"firewall_policy"`
}

// FirewallRuleOptions is used in conjunction with CreateFirewallRule and
// UpdateFirewallRule to create and update firewall rules.
type FirewallRuleOptions struct {
	ID              string  `json:"-"`
	FirewallPolicy  *string `json:"firewall_policy,omitempty"`
	Protocol        *string `json:"protocol,omitempty"`
	Source          *string `json:"source,omitempty"`
	SourcePort      *string `json:"source_port,omitempty"`
	Destination     *string `json:"destination,omitempty"`
	DestinationPort *string `json:"destination_port,omitempty"`
	IcmpTypeName    *string `json:"icmp_type_name,omitempty"`
	Description     *string `json:"description,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c FirewallRule) APIPath() string {
	return "firewall_rules"
}

// PostPath returns the relative URL path to POST an object
func (c FirewallRule) PostPath(from *FirewallRuleOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c FirewallRule) PutPath(from *FirewallRuleOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c FirewallRule) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// FetchID returns the ID field from the object
func (c FirewallRule) FetchID() string {
	return c.ID
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c FirewallRuleOptions) OptionID() string {
	return c.ID
}

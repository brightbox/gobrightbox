package brightbox

import (
	"time"
)

// FirewallRule represents a firewall rule.
// https://api.gb1.brightbox.com/1.0/#firewall_rule
type FirewallRule struct {
	Resource
	Source          string
	SourcePort      string `json:"source_port"`
	Destination     string
	DestinationPort string `json:"destination_port"`
	Protocol        string
	IcmpTypeName    string     `json:"icmp_type_name"`
	CreatedAt       time.Time `json:"created_at"`
	Description     *string
	FirewallPolicy  FirewallPolicy `json:"firewall_policy"`
}

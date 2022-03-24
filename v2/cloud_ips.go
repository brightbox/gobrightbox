package brightbox

import (
	"context"
	"path"

	"github.com/brightbox/gobrightbox/v2/status/cloudip"
)

//go:generate ./generate_status_enum cloudip mapped reserved unmapped

// CloudIP represents a Cloud IP
// https://api.gb1.brightbox.com/1.0/#cloud_ip
type CloudIP struct {
	ResourceRef
	ID              string
	Name            string
	PublicIP        string `json:"public_ip"`
	PublicIPv4      string `json:"public_ipv4"`
	PublicIPv6      string `json:"public_ipv6"`
	Status          cloudip.Status
	ReverseDNS      string `json:"reverse_dns"`
	Fqdn            string
	Mode            string
	Account         *Account
	Interface       *Interface
	Server          *Server
	ServerGroup     *ServerGroup     `json:"server_group"`
	PortTranslators []PortTranslator `json:"port_translators"`
	LoadBalancer    *LoadBalancer    `json:"load_balancer"`
	DatabaseServer  *DatabaseServer  `json:"database_server"`
}

// PortTranslator represents a port translator on a Cloud IP
type PortTranslator struct {
	Incoming int    `json:"incoming"`
	Outgoing int    `json:"outgoing"`
	Protocol string `json:"protocol"`
}

// CloudIPOptions is used in conjunction with CreateCloudIP and UpdateCloudIP to
// create and update cloud IPs.
type CloudIPOptions struct {
	ID              string           `json:"-"`
	ReverseDNS      *string          `json:"reverse_dns,omitempty"`
	Name            *string          `json:"name,omitempty"`
	PortTranslators []PortTranslator `json:"port_translators,omitempty"`
}

// MapCloudIP issues a request to map the cloud ip to the destination. The
// destination can be an identifier of any resource capable of receiving a Cloud
// IP, such as a server interface, a load balancer, or a cloud sql instace.
func (c *Client) MapCloudIP(ctx context.Context, identifier string, destination string) (*CloudIP, error) {
	return APIPost[CloudIP](
		ctx,
		c,
		path.Join(CloudIPAPIPath, identifier, "map"),
		map[string]string{"destination": destination},
	)
}

// UnMapCloudIP issues a request to unmap the cloud ip.
func (c *Client) UnMapCloudIP(ctx context.Context, identifier string) (*CloudIP, error) {
	return APIPost[CloudIP](
		ctx,
		c,
		path.Join(CloudIPAPIPath, identifier, "unmap"),
		nil,
	)
}

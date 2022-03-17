package brightbox

import (
	"context"
	"path"
	"time"

	"github.com/brightbox/gobrightbox/v2/status/balancingpolicy"
	"github.com/brightbox/gobrightbox/v2/status/healthchecktype"
	"github.com/brightbox/gobrightbox/v2/status/listenerprotocol"
	"github.com/brightbox/gobrightbox/v2/status/loadbalancer"
	"github.com/brightbox/gobrightbox/v2/status/proxyprotocol"
)

//go:generate ./generate_status_enum loadbalancer creating active deleting deleted failing failed
//go:generate ./generate_status_enum proxyprotocol v1 v2 v2-ssl v2-ssl-cn
//go:generate ./generate_status_enum balancingpolicy least-connections round-robin source-address
//go:generate ./generate_status_enum healthchecktype tcp http
//go:generate ./generate_status_enum listenerprotocol tcp http https

// LoadBalancer represents a Load Balancer
// https://api.gb1.brightbox.com/1.0/#load_balancer
type LoadBalancer struct {
	ResourceRef
	ID                string
	Name              string
	Status            loadbalancer.Status
	CreatedAt         *time.Time `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Locked            bool
	HTTPSRedirect     bool   `json:"https_redirect"`
	SslMinimumVersion string `json:"ssl_minimum_version"`
	Account           Account
	Nodes             []Server
	CloudIPs          []CloudIP `json:"cloud_ips"`
	Policy            string
	BufferSize        int `json:"buffer_size"`
	Listeners         []LoadBalancerListener
	Healthcheck       LoadBalancerHealthcheck
	Certificate       *LoadBalancerCertificate
	Acme              *LoadBalancerAcme
}

// LoadBalancerCertificate represents a certificate on a LoadBalancer
type LoadBalancerCertificate struct {
	ExpiresAt time.Time `json:"expires_at"`
	ValidFrom time.Time `json:"valid_from"`
	SslV3     bool      `json:"sslv3"`
	Issuer    string    `json:"issuer"`
	Subject   string    `json:"subject"`
}

// LoadBalancerAcme represents an ACME object on a LoadBalancer
type LoadBalancerAcme struct {
	Certificate *LoadBalancerAcmeCertificate `json:"certificate"`
	Domains     []LoadBalancerAcmeDomain     `json:"domains"`
}

// LoadBalancerAcmeCertificate represents an ACME issued certificate on
// a LoadBalancer
type LoadBalancerAcmeCertificate struct {
	Fingerprint string    `json:"fingerprint"`
	ExpiresAt   time.Time `json:"expires_at"`
	IssuedAt    time.Time `json:"issued_at"`
}

// LoadBalancerAcmeDomain represents a domain for which ACME support
// has been requested
type LoadBalancerAcmeDomain struct {
	Identifier  string `json:"identifier"`
	Status      string `json:"status"`
	LastMessage string `json:"last_message"`
}

// LoadBalancerHealthcheck represents a health check on a LoadBalancer
type LoadBalancerHealthcheck struct {
	Type          healthchecktype.Status `json:"type"`
	Port          int                    `json:"port"`
	Request       string                 `json:"request,omitempty"`
	Interval      int                    `json:"interval,omitempty"`
	Timeout       int                    `json:"timeout,omitempty"`
	ThresholdUp   int                    `json:"threshold_up,omitempty"`
	ThresholdDown int                    `json:"threshold_down,omitempty"`
}

// LoadBalancerListener represents a listener on a LoadBalancer
type LoadBalancerListener struct {
	Protocol      listenerprotocol.Status `json:"protocol,omitempty"`
	In            int                     `json:"in,omitempty"`
	Out           int                     `json:"out,omitempty"`
	Timeout       int                     `json:"timeout,omitempty"`
	ProxyProtocol proxyprotocol.Status    `json:"proxy_protocol,omitempty"`
}

// LoadBalancerOptions is used in conjunction with CreateLoadBalancer and
// UpdateLoadBalancer to create and update load balancers
type LoadBalancerOptions struct {
	ID                    string                   `json:"-"`
	Name                  *string                  `json:"name,omitempty"`
	Nodes                 []LoadBalancerNode       `json:"nodes,omitempty"`
	Policy                balancingpolicy.Status   `json:"policy,omitempty"`
	Listeners             []LoadBalancerListener   `json:"listeners,omitempty"`
	Healthcheck           *LoadBalancerHealthcheck `json:"healthcheck,omitempty"`
	Domains               *[]string                `json:"domains,omitempty"`
	CertificatePem        *string                  `json:"certificate_pem,omitempty"`
	CertificatePrivateKey *string                  `json:"certificate_private_key,omitempty"`
	SslMinimumVersion     *string                  `json:"ssl_minimum_version,omitempty"`
	HTTPSRedirect         *bool                    `json:"https_redirect,omitempty"`
}

// LoadBalancerNode is used in conjunction with LoadBalancerOptions,
// AddNodesToLoadBalancer, RemoveNodesFromLoadBalancer to specify a list of
// servers to use as load balancer nodes. The Node parameter should be a server
// identifier.
type LoadBalancerNode struct {
	Node string `json:"node,omitempty"`
}

// AddNodesToLoadBalancer adds nodes to an existing load balancer.
func (c *Client) AddNodesToLoadBalancer(ctx context.Context, identifier string, nodes []LoadBalancerNode) (*LoadBalancer, error) {
	return APIPost[LoadBalancer](
		ctx,
		c,
		path.Join(LoadBalancerAPIPath, identifier, "add_nodes"),
		nodes,
	)

}

// RemoveNodesFromLoadBalancer removes nodes from an existing load balancer.
func (c *Client) RemoveNodesFromLoadBalancer(ctx context.Context, identifier string, nodes []LoadBalancerNode) (*LoadBalancer, error) {
	return APIPost[LoadBalancer](
		ctx,
		c,
		path.Join(LoadBalancerAPIPath, identifier, "remove_nodes"),
		nodes,
	)
}

// AddListenersToLoadBalancer adds listeners to an existing load balancer.
func (c *Client) AddListenersToLoadBalancer(ctx context.Context, identifier string, listeners []LoadBalancerListener) (*LoadBalancer, error) {
	return APIPost[LoadBalancer](
		ctx,
		c,
		path.Join(LoadBalancerAPIPath, identifier, "add_listeners"),
		listeners,
	)

}

// RemoveListenersFromLoadBalancer removes listeners from an existing load balancer.
func (c *Client) RemoveListenersFromLoadBalancer(ctx context.Context, identifier string, listeners []LoadBalancerListener) (*LoadBalancer, error) {
	return APIPost[LoadBalancer](
		ctx,
		c,
		path.Join(LoadBalancerAPIPath, identifier, "remove_listeners"),
		listeners,
	)
}

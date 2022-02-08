package brightbox

import (
	"time"
)

// LoadBalancer represents a Load Balancer
// https://api.gb1.brightbox.com/1.0/#load_balancer
type LoadBalancer struct {
	ResourceRef
	ID                string
	Name              string
	Status            string
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
	Type          string `json:"type"`
	Port          int    `json:"port"`
	Request       string `json:"request,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	ThresholdUp   int    `json:"threshold_up,omitempty"`
	ThresholdDown int    `json:"threshold_down,omitempty"`
}

// LoadBalancerListener represents a listener on a LoadBalancer
type LoadBalancerListener struct {
	Protocol      string `json:"protocol,omitempty"`
	In            int    `json:"in,omitempty"`
	Out           int    `json:"out,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	ProxyProtocol string `json:"proxy_protocol,omitempty"`
}

// LoadBalancerOptions is used in conjunction with CreateLoadBalancer and
// UpdateLoadBalancer to create and update load balancers
type LoadBalancerOptions struct {
	ID                    string                   `json:"-"`
	Name                  *string                  `json:"name,omitempty"`
	Nodes                 []LoadBalancerNode       `json:"nodes,omitempty"`
	Policy                *string                  `json:"policy,omitempty"`
	BufferSize            *int                     `json:"buffer_size,omitempty"`
	Listeners             []LoadBalancerListener   `json:"listeners,omitempty"`
	Healthcheck           *LoadBalancerHealthcheck `json:"healthcheck,omitempty"`
	Domains               *[]string                `json:"domains,omitempty"`
	CertificatePem        *string                  `json:"certificate_pem,omitempty"`
	CertificatePrivateKey *string                  `json:"certificate_private_key,omitempty"`
	SslMinimumVersion     *string                  `json:"ssl_minimum_version,omitempty"`
	SslV3                 *bool                    `json:"sslv3,omitempty"`
	HTTPSRedirect         *bool                    `json:"https_redirect,omitempty"`
}

// LoadBalancerNode is used in conjunction with LoadBalancerOptions,
// AddNodesToLoadBalancer, RemoveNodesFromLoadBalancer to specify a list of
// servers to use as load balancer nodes. The Node parameter should be a server
// identifier.
type LoadBalancerNode struct {
	Node string `json:"node"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c LoadBalancer) APIPath() string {
	return "load_balancers"
}

// FetchID returns the ID field from the object
func (c LoadBalancer) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c LoadBalancer) PostPath(from *LoadBalancerOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c LoadBalancer) PutPath(from *LoadBalancerOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c LoadBalancer) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c LoadBalancerOptions) OptionID() string {
	return c.ID
}

// LockID returns the path to a lockable object
func (c LoadBalancer) LockID() string {
	return c.APIPath() + "/" + c.FetchID()
}

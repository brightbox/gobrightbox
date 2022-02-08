package brightbox

// CloudIP represents a Cloud IP
// https://api.gb1.brightbox.com/1.0/#cloud_ip
type CloudIP struct {
	ResourceRef
	ID              string
	Name            string
	PublicIP        string `json:"public_ip"`
	PublicIPv4      string `json:"public_ipv4"`
	PublicIPv6      string `json:"public_ipv6"`
	Status          string
	ReverseDNS      string `json:"reverse_dns"`
	Fqdn            string
	Mode            string
	Account         *Account
	Interface       *Interface
	Server          *Server
	ServerGroup     *ServerGroup     `json:"server_group"`
	PortTranslators []PortTranslator `json:"port_translators"`
	LoadBalancer    *LoadBalancer    `json:"load_balancer"`
	//	DatabaseServer  *DatabaseServer `json:"database_server"`
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

// APIPath returns the relative URL path to the collection endpoint
func (c CloudIP) APIPath() string {
	return "cloud_ips"
}

// PostPath returns the relative URL path to POST an object
func (c CloudIP) PostPath(from *CloudIPOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c CloudIP) PutPath(from *CloudIPOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c CloudIP) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// FetchID returns the ID field from the object
func (c CloudIP) FetchID() string {
	return c.ID
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c CloudIPOptions) OptionID() string {
	return c.ID
}

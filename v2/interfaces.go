package brightbox

// Interface represent a server's network interface(s)
// https://api.gb1.brightbox.com/1.0/#interface
type Interface struct {
	ID          string
	MacAddress  string `json:"mac_address"`
	IPv4Address string `json:"ipv4_address"`
	IPv6Address string `json:"ipv6_address"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c Interface) APIPath() string {
	return "interfaces"
}

// FetchID returns the ID field from the object
func (c Interface) FetchID() string {
	return c.ID
}

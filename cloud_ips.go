package brightbox

import (
	"fmt"
)

type CloudIP struct {
	Resource
	Name            string
	PublicIP        string `json:"public_ip"`
	Status          string
	ReverseDns      string           `json:"reverse_dns"`
	PortTranslators []PortTranslator `json:"port_translators"`
	Account         Account
	Fqdn            string
	Interface       *ServerInterface
	Server          *Server
	ServerGroup     *ServerGroup    `json:"server_group"`
	LoadBalancer    *LoadBalancer   `json:"load_balancer"`
	DatabaseServer  *DatabaseServer `json:"database_server"`
}

type PortTranslator struct {
	Incoming int
	Outgoing int
	Protocol string
}

func (c *Client) CloudIPs() (*[]CloudIP, error) {
	cloudips := new([]CloudIP)
	_, err := c.MakeApiRequest("GET", "/1.0/cloud_ips", nil, cloudips)
	if err != nil {
		return nil, err
	}
	return cloudips, err
}

func (c *Client) CloudIP(identifier string) (*CloudIP, error) {
	cloudip := new(CloudIP)
	_, err := c.MakeApiRequest("GET", "/1.0/cloud_ips/"+identifier, nil, cloudip)
	if err != nil {
		return nil, err
	}
	return cloudip, err
}

func (c *Client) DestroyCloudIP(identifier string) error {
	_, err := c.MakeApiRequest("DELETE", "/1.0/cloud_ips/"+identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// MapCloudIP issues a request to map the cloud ip to the destination. The
// destination can be an identifier of any resource capable of receiving a Cloud
// IP, such as a server interface, a load balancer, or a cloud sql instace.
//
// To map a Cloud IP to a server, first lookup the server to get it's interface
// identifier (or use the MapCloudIPtoServer convenience method)
func (c *Client) MapCloudIP(identifier string, destination string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/cloud_ips/"+identifier+"/map",
		map[string]string{"destination": destination}, nil)
	if err != nil {
		return err
	}
	return nil
}

// Convenience method to map a Cloud IP to a server. First looks up the server
// to get the network interface id. Uses the first interface found.
func (c *Client) MapCloudIPtoServer(identifier string, serverid string) error {
	server, err := c.Server(serverid)
	if err != nil {
		return err
	}
	if len(server.Interfaces) == 0 {
		return fmt.Errorf("Server %s has no interfaces to map cloud ip %s to", server.Id, identifier)
	}
	destination := server.Interfaces[0].Id
	err = c.MapCloudIP(identifier, destination)
	if err != nil {
		return err
	}
	return nil
}

// UnMapCloudIP issues a request to unmap the cloud ip.
func (c *Client) UnMapCloudIP(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/cloud_ips/"+identifier+"/unmap", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

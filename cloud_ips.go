package brightbox

type CloudIP struct {
	Resource
	Name            string
	PublicIP        string `json:"public_ip"`
	Status          string
	ReverseDns      string           `json:"reverse_dns"`
	PortTranslators []PortTranslator `json:"port_translators"`
	Account         Account
	Interface       *ServerInterface
	Server          *Server
	//LoadBalancer
	//DatabaseServer
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

package brightbox

import (
	"regexp"
	"time"
)

type Server struct {
	Resource
	Name                string
	Status              string
	Locked              bool
	Hostname            string
	Fqdn                string
	CreatedAt           *time.Time `json:"created_at"`
	DeletedAt           *time.Time `json:"deleted_at"`
	ServerType          ServerType `json:"server_type"`
	CompatabilityMode   bool       `json:"compatibility_mode"`
	Zone                Zone
	Image               Image
	CloudIPs            []CloudIP `json:"cloud_ips"`
	Interfaces          []ServerInterface
	Snapshots           []Image
	ServerGroups        []ServerGroup `json:"server_groups"`
	ConsoleToken        string        `json:"console_token"`
	ConsoleUrl          string        `json:"console_url"`
	ConsoleTokenExpires *time.Time    `json:"console_token_expires"`
}

type CreateServerOptions struct {
	Identifier   string   `json:"-"`
	Image        string   `json:"image"`
	Name         string   `json:"name,omitempty"`
	ServerType   string   `json:"server_type,omitempty"`
	Zone         string   `json:"zone,omitempty"`
	UserData     string   `json:"user_data,omitempty"`
	ServerGroups []string `json:"server_groups,omitempty"`
}

type ServerInterface struct {
	Resource
	MacAddress  string `json:"mac_address"`
	IPv4Address string `json:"ipv4_address"`
	IPv6Address string `json:"ipv6_address"`
}

func (c *Client) Servers() (*[]Server, error) {
	servers := new([]Server)
	_, err := c.MakeApiRequest("GET", "/1.0/servers", nil, servers)
	if err != nil {
		return nil, err
	}
	return servers, err
}

func (c *Client) Server(identifier string) (*Server, error) {
	server := new(Server)
	_, err := c.MakeApiRequest("GET", "/1.0/servers/"+identifier, nil, server)
	if err != nil {
		return nil, err
	}
	return server, err
}

func (c *Client) CreateServer(newServer *CreateServerOptions) (*Server, error) {
	server := new(Server)
	_, err := c.MakeApiRequest("POST", "/1.0/servers", newServer, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (c *Client) DestroyServer(identifier string) error {
	_, err := c.MakeApiRequest("DELETE", "/1.0/servers/"+identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StopServer(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/stop", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StartServer(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/start", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RebootServer(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/reboot", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ResetServer(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/reset", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ShutdownServer(identifier string) error {
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/shutdown", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) LockServer(identifier string) error {
	_, err := c.MakeApiRequest("PUT", "/1.0/servers/"+identifier+"/lock_resource", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UnlockServer(identifier string) error {
	_, err := c.MakeApiRequest("PUT", "/1.0/servers/"+identifier+"/unlock_resource", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SnapshotServer(identifier string) (*Image, error) {
	res, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/snapshot", nil, nil)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile("img-.....")
	imageId := re.FindString(res.Header.Get("Link"))
	if imageId != "" {
		img := new(Image)
		img.Id = imageId
		return img, nil
	}
	return nil, nil
}

func (c *Client) ActivateConsoleForServer(identifier string) (*Server, error) {
	server := new(Server)
	_, err := c.MakeApiRequest("POST", "/1.0/servers/"+identifier+"/activate_console", nil, server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

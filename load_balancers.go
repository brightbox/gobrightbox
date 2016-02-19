package brightbox

import (
	"time"
)

type LoadBalancer struct {
	Resource
	Name       string
	Status     string
	CreatedAt  *time.Time `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Locked     bool
	Account    Account
	Nodes      []Server
	CloudIPs   []CloudIP `json:"cloud_ips"`
	Policy     string
	BufferSize int `json:"buffer_size"`
	// Certificate FIXME
	// Listeners FIXME
	// Healthcheck FIXME
}

func (c *Client) LoadBalancers() ([]LoadBalancer, error) {
	var lbs []LoadBalancer
	_, err := c.MakeApiRequest("GET", "/1.0/load_balancers", nil, &lbs)
	if err != nil {
		return nil, err
	}
	return lbs, err
}

func (c *Client) LoadBalancer(identifier string) (*LoadBalancer, error) {
	lb := new(LoadBalancer)
	_, err := c.MakeApiRequest("GET", "/1.0/load_balancers/"+identifier, nil, lb)
	if err != nil {
		return nil, err
	}
	return lb, err
}

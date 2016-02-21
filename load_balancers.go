package brightbox

import (
	"time"
)

type LoadBalancer struct {
	Resource
	Name        string
	Status      string
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Locked      bool
	Account     Account
	Nodes       []Server
	CloudIPs    []CloudIP `json:"cloud_ips"`
	Policy      string
	BufferSize  int `json:"buffer_size"`
	Listeners   []LoadBalancerListener
	Healthcheck LoadBalancerHealthCheck
	Certificate *LoadBalancerCertificate
}

type LoadBalancerCertificate struct {
	ExpiresAt time.Time `json:"expires_at"`
	ValidFrom time.Time `json:"valid_from"`
	SslV3     bool      `json:"sslv3"`
	Issuer    string    `json:"issuer"`
	Subject   string    `json:"subject"`
}

type LoadBalancerHealthCheck struct {
	Type          string `json:"type"`
	Port          int    `json:"port"`
	Request       string `json:"request,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	ThresholdUp   int    `json:"threshold_up,omitempty"`
	ThresholdDown int    `json:"threshold_down,omitempty"`
}

type LoadBalancerListener struct {
	Protocol string `json:"protocol"`
	In       int    `json:"in"`
	Out      int    `json:"out"`
	Timeout  int    `json:"timeout,omitempty"`
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

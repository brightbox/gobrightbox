package gobrightbox

import (
	"time"
)

// APIClient represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type APIClient struct {
	ID               string
	Name             string
	Description      string
	Secret           string
	PermissionsGroup string     `json:"permissions_group"`
	RevokedAt        *time.Time `json:"revoked_at"`
	Account          Account
}

// APIClientOptions is used in conjunction with CreateAPIClient and
// UpdateAPIClient to create and update api clients
type APIClientOptions struct {
	ID               string  `json:"-"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	PermissionsGroup *string `json:"permissions_group,omitempty"`
}

// APIClients retrieves a list of all API clients
func (c *Client) APIClients() ([]APIClient, error) {
	var apiClients []APIClient
	_, err := c.MakeAPIRequest("GET", "/1.0/api_clients", nil, &apiClients)
	if err != nil {
		return nil, err
	}
	return apiClients, err
}

// APIClient retrieves a detailed view of one API client
func (c *Client) APIClient(identifier string) (*APIClient, error) {
	apiClient := new(APIClient)
	_, err := c.MakeAPIRequest("GET", "/1.0/api_clients/"+identifier, nil, apiClient)
	if err != nil {
		return nil, err
	}
	return apiClient, err
}

// CreateAPIClient creates a new API client.
//
// It takes a APIClientOptions struct for specifying name and other
// attributes. Not all attributes can be specified at create time
// (such as ID, which is allocated for you)
func (c *Client) CreateAPIClient(options *APIClientOptions) (*APIClient, error) {
	ac := new(APIClient)
	_, err := c.MakeAPIRequest("POST", "/1.0/api_clients", options, &ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

// UpdateAPIClient updates an existing api client.
//
// It takes a APIClientOptions struct for specifying ID, name and other
// attributes. Not all attributes can be specified at update time.
func (c *Client) UpdateAPIClient(options *APIClientOptions) (*APIClient, error) {
	ac := new(APIClient)
	_, err := c.MakeAPIRequest("PUT", "/1.0/api_clients/"+options.ID, options, &ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

// DestroyAPIClient issues a request to deletes an existing api client
func (c *Client) DestroyAPIClient(identifier string) error {
	_, err := c.MakeAPIRequest("DELETE", "/1.0/api_clients/"+identifier, nil, nil)
	return err
}

// ResetSecretForAPIClient requests a snapshot of an existing api client
func (c *Client) ResetSecretForAPIClient(identifier string) (*APIClient, error) {
	ac := new(APIClient)
	_, err := c.MakeAPIRequest("POST", "/1.0/api_clients/"+identifier+"/reset_secret", nil, &ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

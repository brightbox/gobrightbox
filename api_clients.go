package brightbox

type ApiClient struct {
	Resource
	Name             string
	Description      string
	Secret           string
	PermissionsGroup string
	Account          Account
}

func (c *Client) ApiClients() (*[]ApiClient, error) {
	apiClients := make([]ApiClient, 1)
	_, err := c.MakeApiRequest("GET", "/1.0/api_clients", nil, apiClients)
	if err != nil {
		return nil, err
	}
	return &apiClients, err
}

func (c *Client) ApiClient(identifier string) (*ApiClient, error) {
	apiClient := new(ApiClient)
	_, err := c.MakeApiRequest("GET", "/1.0/api_clients/"+identifier, nil, apiClient)
	if err != nil {
		return nil, err
	}
	return apiClient, err
}

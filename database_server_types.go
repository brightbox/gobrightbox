package gobrightbox

// DatabaseServerType represents a database server type
// https://api.gb1.brightbox.com/1.0/#database_type
type DatabaseServerType struct {
	ID          string
	Name        string
	Description string
	DiskSize    int `json:"disk_size"`
	RAM         int
}

// DatabaseServerTypes retrieves a list of Database Types
func (c *Client) DatabaseServerTypes() ([]DatabaseServerType, error) {
	var databaseservertypes []DatabaseServerType
	_, err := c.MakeAPIRequest("GET", "/1.0/database_types", nil, &databaseservertypes)
	if err != nil {
		return nil, err
	}
	return databaseservertypes, err
}

// DatabaseServerType retrieves a detailed view of one Database Type
func (c *Client) DatabaseServerType(identifier string) (*DatabaseServerType, error) {
	databaseservertype := new(DatabaseServerType)
	_, err := c.MakeAPIRequest("GET", "/1.0/database_types/"+identifier, nil, databaseservertype)
	if err != nil {
		return nil, err
	}
	return databaseservertype, err
}

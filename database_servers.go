package brightbox

import (
	"time"
)

// DatabaseServer represents a database server.
// https://api.gb1.brightbox.com/1.0/#database_server
type DatabaseServer struct {
	Id                 string
	Name               string
	Description        string
	Status             string
	Account            Account
	DatabaseEngine     string     `json:"database_engine"`
	DatabaseVersion    string     `json:"database_version"`
	AdminUsername      string     `json:"admin_username"`
	AdminPassword      string     `json:"admin_password"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	AllowAccess        []string `json:"allow_access"`
	MaintenanceWeekday int `json:"maintenance_weekday"`
	MaintenanceHour    int `json:"maintenance_hour"`
	Locked             bool
	CloudIPs           []CloudIP `json:"cloud_ips"`
	Zone               Zone
	DatabaseServerType DatabaseServerType `json:"database_server_type"`
}

// DatabaseServerType represents a database server type
// https://api.gb1.brightbox.com/1.0/#database_type
type DatabaseServerType struct {
	Id          string
	Name        string
	Description string
	DiskSize    int `json:"disk_size"`
	Ram         int
}

// DatabaseServerOptions is used in conjunction with CreateDatabaseServer and
// UpdateDatabaseServer to create and update database servers.
type DatabaseServerOptions struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Engine      *string   `json:"engine,omitempty"`
	Version     *string   `json:"version,omitempty"`
	AllowAccess *[]string `json:"allow_access,omitempty"`
	Snapshot    *string   `json:"snapshot,omitempty"`
	Zone        *string   `json:"snapshot,omitempty"`
}

// DatabaseServers retrieves a list of all database servers
func (c *Client) DatabaseServers() ([]DatabaseServer, error) {
	var dbs []DatabaseServer
	_, err := c.MakeApiRequest("GET", "/1.0/database_servers", nil, &dbs)
	if err != nil {
		return nil, err
	}
	return dbs, err
}

// DatabaseServer retrieves a detailed view of one database server
func (c *Client) DatabaseServer(identifier string) (*DatabaseServer, error) {
	dbs := new(DatabaseServer)
	_, err := c.MakeApiRequest("GET", "/1.0/database_servers/"+identifier, nil, dbs)
	if err != nil {
		return nil, err
	}
	return dbs, err
}

// CreateDatabaseServer creates a new database server.
//
// It takes a DatabaseServerOptions struct for specifying name and other
// attributes. Not all attributes can be specified at create time
// (such as Id, which is allocated for you)
func (c *Client) CreateDatabaseServer(options *DatabaseServerOptions) (*DatabaseServer, error) {
	dbs := new(DatabaseServer)
	_, err := c.MakeApiRequest("POST", "/1.0/database_servers", options, &dbs)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

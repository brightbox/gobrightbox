package brightbox

import (
	"time"
)

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
	AllowAccess        []string
	MaintenanceWeekday int `json:"maintenance_weekday"`
	MaintenanceHour    int `json:"maintenance_hour"`
	Locked             bool
	CloudIPs           []CloudIP `json:"cloud_ips"`
	Zone               Zone
	DatabaseServerType DatabaseServerType `json:"database_server_type"`
}

type DatabaseServerType struct {
	Name        string
	Description string
	DiskSize    int `json:"disk_size"`
	Ram         int
}

func (c *Client) DatabaseServers() ([]DatabaseServer, error) {
	var dbs []DatabaseServer
	_, err := c.MakeApiRequest("GET", "/1.0/database_servers", nil, &dbs)
	if err != nil {
		return nil, err
	}
	return dbs, err
}

func (c *Client) DatabaseServer(identifier string) (*DatabaseServer, error) {
	dbs := new(DatabaseServer)
	_, err := c.MakeApiRequest("GET", "/1.0/database_servers/"+identifier, nil, dbs)
	if err != nil {
		return nil, err
	}
	return dbs, err
}

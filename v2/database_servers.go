package brightbox

import (
	"time"
)

// DatabaseServer represents a database server.
// https://api.gb1.brightbox.com/1.0/#database_server
type DatabaseServer struct {
	ResourceRef
	ID                      string
	Name                    string
	Description             string
	Status                  string
	DatabaseEngine          string     `json:"database_engine"`
	DatabaseVersion         string     `json:"database_version"`
	AdminUsername           string     `json:"admin_username"`
	AdminPassword           string     `json:"admin_password"`
	CreatedAt               *time.Time `json:"created_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	SnapshotsSchedule       *string    `json:"snapshots_schedule"`
	SnapshotsScheduleNextAt *time.Time `json:"snapshots_schedule_next_at"`
	AllowAccess             []string   `json:"allow_access"`
	MaintenanceWeekday      int        `json:"maintenance_weekday"`
	MaintenanceHour         int        `json:"maintenance_hour"`
	Locked                  bool
	Account                 *Account
	Zone                    *Zone
	DatabaseServerType      *DatabaseServerType `json:"database_server_type"`
	CloudIPs                []CloudIP           `json:"cloud_ips"`
}

// DatabaseServerOptions is used in conjunction with CreateDatabaseServer and
// UpdateDatabaseServer to create and update database servers.
type DatabaseServerOptions struct {
	ID                 string   `json:"-"`
	Name               *string  `json:"name,omitempty"`
	Description        *string  `json:"description,omitempty"`
	Engine             string   `json:"engine,omitempty"`
	Version            string   `json:"version,omitempty"`
	AllowAccess        []string `json:"allow_access,omitempty"`
	Snapshot           string   `json:"snapshot,omitempty"`
	Zone               string   `json:"zone,omitempty"`
	DatabaseType       string   `json:"database_type,omitempty"`
	MaintenanceWeekday *int     `json:"maintenance_weekday,omitempty"`
	MaintenanceHour    *int     `json:"maintenance_hour,omitempty"`
	SnapshotsSchedule  *string  `json:"snapshots_schedule,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c DatabaseServer) APIPath() string {
	return "database_servers"
}

// FetchID returns the ID field from the object
func (c DatabaseServer) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c DatabaseServer) PostPath(from *DatabaseServerOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c DatabaseServer) PutPath(from *DatabaseServerOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c DatabaseServer) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c DatabaseServerOptions) OptionID() string {
	return c.ID
}

// LockID returns the path to a lockable object
func (c DatabaseServer) LockID() string {
	return c.APIPath() + "/" + c.FetchID()
}

// ResetPasswordPath returns the relative URL path to reset the password
func (c DatabaseServer) ResetPasswordPath() string {
	return c.APIPath() + "/" + c.FetchID() + "/reset_password"
}

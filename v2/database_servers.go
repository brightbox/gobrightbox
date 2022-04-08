package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/databaseserver"
)

//go:generate ./generate_status_enum databaseserver creating active deleting deleted failing failed

// DatabaseServer represents a database server.
// https://api.gb1.brightbox.com/1.0/#database_server
type DatabaseServer struct {
	ResourceRef
	ID                      string
	Name                    string
	Description             string
	Status                  databaseserver.Status
	DatabaseEngine          string   `json:"database_engine"`
	DatabaseVersion         string   `json:"database_version"`
	AdminUsername           string   `json:"admin_username"`
	AdminPassword           string   `json:"admin_password"`
	SnapshotsSchedule       *string  `json:"snapshots_schedule"`
	AllowAccess             []string `json:"allow_access"`
	MaintenanceWeekday      int      `json:"maintenance_weekday"`
	MaintenanceHour         int      `json:"maintenance_hour"`
	Locked                  bool
	CreatedAt               *time.Time `json:"created_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	SnapshotsScheduleNextAt *time.Time `json:"snapshots_schedule_next_at"`
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

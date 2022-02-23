package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/databasesnapshot"
)

//go:generate ./generate_status_enum databasesnapshot creating available deleted failed

// DatabaseSnapshot represents a snapshot of a database server.
// https://api.gb1.brightbox.com/1.0/#databaseSnapshot
type DatabaseSnapshot struct {
	ResourceRef
	ID              string
	Name            string
	Description     string
	Status          databasesnapshot.Status
	DatabaseEngine  string `json:"database_engine"`
	DatabaseVersion string `json:"database_version"`
	Source          string
	SourceTrigger   string `json:"source_trigger"`
	Size            int
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Locked          bool
	Account         Account
}

// DatabaseSnapshotOptions is used to update snapshots
type DatabaseSnapshotOptions struct {
	ID          string  `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

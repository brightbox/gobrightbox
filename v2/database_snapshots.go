package brightbox

import (
	"time"
)

// DatabaseSnapshot represents a snapshot of a database server.
// https://api.gb1.brightbox.com/1.0/#databaseSnapshot
type DatabaseSnapshot struct {
	ResourceRef
	ID              string
	Name            string
	Description     string
	Status          string
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

// APIPath returns the relative URL path to the collection endpoint
func (c DatabaseSnapshot) APIPath() string {
	return "database_snapshots"
}

// FetchID returns the ID field from the object
func (c DatabaseSnapshot) FetchID() string {
	return c.ID
}

// PutPath returns the relative URL path to PUT an object
func (c DatabaseSnapshot) PutPath(from *DatabaseSnapshotOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c DatabaseSnapshot) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c DatabaseSnapshotOptions) OptionID() string {
	return c.ID
}

// LockID returns the path to a lockable object
func (c DatabaseSnapshot) LockID() string {
	return c.APIPath() + "/" + c.FetchID()
}

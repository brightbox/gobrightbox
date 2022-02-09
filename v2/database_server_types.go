package brightbox

// DatabaseServerType represents a database server type
// https://api.gb1.brightbox.com/1.0/#database_type
type DatabaseServerType struct {
	ResourceRef
	ID          string
	Name        string
	Description string
	RAM         int
	DiskSize    int `json:"disk_size"`
	Default     bool
}

// APIPath returns the relative URL path to the collection endpoint
func (c DatabaseServerType) APIPath() string {
	return "database_types"
}

// FetchID returns the ID field from the object
func (c DatabaseServerType) FetchID() string {
	return c.ID
}

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

package brightbox

// ServerType represents a Server Type
// https://api.gb1.brightbox.com/1.0/#server_type
type ServerType struct {
	ResourceRef
	ID       string
	Name     string
	Status   string
	Handle   string
	Cores    int
	RAM      int
	DiskSize int `json:"disk_size"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c ServerType) APIPath() string {
	return "server_types"
}

// FetchID returns the ID field from the object
func (c ServerType) FetchID() string {
	return c.ID
}

// HandleString returns the Handle field from a ServerType object
func (c ServerType) HandleString() string {
	return c.Handle
}

package brightbox

// Volume represents a Brightbox Volume
// https://api.gb1.brightbox.com/1.0/#volume
type Volume struct {
	ResourceRef
	ID          string
	Name        string
	Status      string
	Description string
	Encrypted   bool
	Size        int
	StorageType string `json:"storage_type"`
	Account     *Account
	Image       *Image
}

// APIPath returns the relative URL path to the collection endpoint
func (c Volume) APIPath() string {
	return "volumes"
}

// FetchID returns the ID field from the object
func (c Volume) FetchID() string {
	return c.ID
}

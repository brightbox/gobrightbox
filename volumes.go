package gobrightbox

// Volume represents a Brightbox Volume
// https://api.gb1.brightbox.com/1.0/#volume
type Volume struct {
	ID          string
	Name        string
	Status      string
	Description string
	Encrypted   bool
	Size        int
	StorageType string `json:"storage_type"`
	Server      *Server
	Account     *Account
	Image       *Image
}

// VolumeOptions is used to create and update volumes
// create and update servers.
type VolumeOptions struct {
	ID    string  `json:"-"`
	Size  *int    `json:"size,omitempty"`
	Image *string `json:"image,omitempty"`
}

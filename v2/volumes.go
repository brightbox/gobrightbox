package brightbox

import (
	"github.com/brightbox/gobrightbox/v2/status/volume"
)

//go:generate ./generate_status_enum volume creating attached detached deleting deleted failed

// Volume represents a Brightbox Volume
// https://api.gb1.brightbox.com/1.0/#volume
type Volume struct {
	ResourceRef
	ID          string
	Name        string
	Status      volume.Status
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

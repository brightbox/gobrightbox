package brightbox

import (
	"github.com/brightbox/gobrightbox/v2/status/servertype"
)

//go:generate ./generate_status_enum servertype experimental available deprecated

// ServerType represents a Server Type
// https://api.gb1.brightbox.com/1.0/#server_type
type ServerType struct {
	ResourceRef
	ID          string
	Name        string
	Status      servertype.Status
	Cores       int
	RAM         int
	Handle      string
	DiskSize    int    `json:"disk_size"`
	StorageType string `json:"storage_type"`
}

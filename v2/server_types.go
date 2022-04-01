package brightbox

import (
	"github.com/brightbox/gobrightbox/v2/status/servertype"
	"github.com/brightbox/gobrightbox/v2/status/storagetype"
)

//go:generate ./generate_status_enum servertype experimental available deprecated
//go:generate ./generate_status_enum storagetype local network

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
	DiskSize    int                `json:"disk_size"`
	StorageType storagetype.Status `json:"storage_type"`
}

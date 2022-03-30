package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/arch"
	"github.com/brightbox/gobrightbox/v2/status/image"
)

//go:generate ./generate_status_enum image creating available deprecated unavailable deleting deleted failed
//go:generate ./generate_status_enum arch x86_64 i686

// Image represents a Machine Image
// https://api.gb1.brightbox.com/1.0/#image
type Image struct {
	ResourceRef
	ID                string
	Name              string
	Username          string
	Status            image.Status
	Locked            bool
	Description       string
	Source            string
	Arch              arch.Status
	CreatedAt         *time.Time `json:"created_at"`
	Official          bool
	Public            bool
	Owner             string
	SourceTrigger     string `json:"source_trigger"`
	SourceType        string `json:"source_type"`
	VirtualSize       int    `json:"virtual_size"`
	DiskSize          int    `json:"disk_size"`
	MinRAM            *int   `json:"min_ram"`
	CompatibilityMode bool   `json:"compatibility_mode"`
	LicenceName       string `json:"licence_name"`
	Ancestor          *Image
}

// ImageOptions is used to create and update machine images
type ImageOptions struct {
	ID                string      `json:"-"`
	Name              *string     `json:"name,omitempty"`
	Username          *string     `json:"username,omitempty"`
	Status            *string     `json:"status,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Source            string      `json:"source,omitempty"`
	Arch              arch.Status `json:"arch,omitempty"`
	Public            *bool       `json:"public,omitempty"`
	CompatibilityMode *bool       `json:"compatibility_mode,omitempty"`
}

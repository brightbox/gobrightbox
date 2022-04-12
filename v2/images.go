package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/arch"
	"github.com/brightbox/gobrightbox/v2/status/image"
	"github.com/brightbox/gobrightbox/v2/status/sourcetrigger"
	"github.com/brightbox/gobrightbox/v2/status/sourcetype"
)

//go:generate ./generate_status_enum image creating available deprecated unavailable deleting deleted failed
//go:generate ./generate_status_enum arch x86_64 i686
//go:generate ./generate_status_enum sourcetrigger manual schedule
//go:generate ./generate_status_enum sourcetype upload snapshot

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
	Official          bool
	Public            bool
	Owner             string
	SourceTrigger     sourcetrigger.Status `json:"source_trigger"`
	SourceType        sourcetype.Status    `json:"source_type"`
	VirtualSize       uint                 `json:"virtual_size"`
	DiskSize          uint                 `json:"disk_size"`
	MinRAM            *uint                `json:"min_ram"`
	CompatibilityMode bool                 `json:"compatibility_mode"`
	LicenceName       string               `json:"licence_name"`
	CreatedAt         *time.Time           `json:"created_at"`
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

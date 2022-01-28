package brightbox

import (
	"time"
)

// Image represents a Machine Image
// https://api.gb1.brightbox.com/1.0/#image
type Image struct {
	ID                string
	Name              string
	Username          string
	Status            string
	Locked            bool
	Description       string
	Source            string
	Arch              string
	CreatedAt         time.Time `json:"created_at"`
	Official          bool
	Public            bool
	Owner             string
	SourceType        string `json:"source_type"`
	VirtualSize       int    `json:"virtual_size"`
	DiskSize          int    `json:"disk_size"`
	CompatibilityMode bool   `json:"compatibility_mode"`
	AncestorID        string `json:"ancestor_id"`
	LicenceName       string `json:"licence_name"`
}

// ImageOptions is used to create and update machine images
type ImageOptions struct {
	ID                string  `json:"-"`
	Name              *string `json:"name,omitempty"`
	Username          *string `json:"username,omitempty"`
	Status            *string `json:"status,omitempty"`
	Description       *string `json:"description,omitempty"`
	Source            *string `json:"source,omitempty"`
	Arch              *string `json:"arch,omitempty"`
	Public            *bool   `json:"public,omitempty"`
	CompatibilityMode *bool   `json:"compatibility_mode,omitempty"`
}

// APIPath returns the relative URL path to the config map collection
func (c Image) APIPath() string {
	return "images"
}

// Extract copies an Image object to a ImageOptions object
func (c Image) Extract() *ImageOptions {
	return &ImageOptions{
		ID:                c.ID,
		Name:              &c.Name,
		Username:          &c.Username,
		Status:            &c.Status,
		Description:       &c.Description,
		Source:            &c.Source,
		Arch:              &c.Arch,
		Public:            &c.Public,
		CompatibilityMode: &c.CompatibilityMode,
	}
}

// FetchID returns the ID field from an ImageOptions object
// ID will be blank for create, and set for update
func (c ImageOptions) FetchID() string {
	return c.ID
}

// LockID returns the path to a lockable Image object
func (c Image) LockID() string {
	return c.APIPath() + "/" + c.ID
}

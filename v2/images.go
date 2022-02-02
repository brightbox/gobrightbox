package brightbox

import (
	"time"
)

// Image represents a Machine Image
// https://api.gb1.brightbox.com/1.0/#image
type Image struct {
	ResourceRef
	ID                string
	Name              string
	Username          string
	Status            string
	Locked            bool
	Description       string
	Source            string
	Arch              string
	CreatedAt         *time.Time `json:"created_at"`
	Official          bool
	Public            bool
	Owner             string
	SourceType        string `json:"source_type"`
	VirtualSize       int    `json:"virtual_size"`
	DiskSize          int    `json:"disk_size"`
	CompatibilityMode bool   `json:"compatibility_mode"`
	LicenceName       string `json:"licence_name"`
	Ancestor          *Image
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

// APIPath returns the relative URL path to the collection endpoint
func (c Image) APIPath() string {
	return "images"
}

// FetchID returns the ID field from the object
func (c Image) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c Image) PostPath(from *ImageOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c Image) PutPath(from *ImageOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c Image) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c ImageOptions) OptionID() string {
	return c.ID
}

// LockID returns the path to a lockable object
func (c Image) LockID() string {
	return c.APIPath() + "/" + c.FetchID()
}

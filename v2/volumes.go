package brightbox

import (
	"context"
	"path"
	"time"

	"github.com/brightbox/gobrightbox/v2/status/volume"
)

//go:generate ./generate_status_enum volume creating attached detached deleting deleted failed

// Volume represents a Brightbox Volume
// https://api.gb1.brightbox.com/1.0/#volume
type Volume struct {
	ResourceRef
	ID               string
	Name             string
	Status           volume.Status
	Description      string
	DeleteWithServer bool `json:"delete_with_server"`
	Encrypted        bool
	Size             uint
	StorageType      string     `json:"storage_type"`
	CreatedAt        *time.Time `json:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	Server           *Server
	Account          *Account
	Image            *Image
}

// VolumeOptions is used to create and update volumes
// create and update servers.
type VolumeOptions struct {
	ID               string  `json:"-"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	Serial           *string `json:"serial,omitempty"`
	Size             uint    `json:"size,omitempty"`
	Image            string  `json:"image,omitempty"`
	Encrypted        *bool   `json:"encrypted,omitempty"`
	DeleteWithServer *bool   `json:"delete_with_server,omitempty"`
	Volume           string  `json:"volume,omitempty"`
}

// VolumeAttachment is used in conjunction with AttachVolume and DetachVolume
type VolumeAttachment struct {
	Server string `json:"server"`
	Boot   bool   `json:"boot"`
}

// VolumeNewSize is used in conjunction with ResizeVolume
// to specify the change in the disk size
type VolumeNewSize struct {
	From uint `json:"from"`
	To   uint `json:"to"`
}

// AttachVolume issues a request to attach the volume to a particular server and
// optionally mark it as the boot volume
func (c *Client) AttachVolume(ctx context.Context, identifier string, attachment VolumeAttachment) (*Volume, error) {
	return APIPost[Volume](
		ctx,
		c,
		path.Join(VolumeAPIPath, identifier, "attach"),
		attachment,
	)
}

// DetachVolume issues a request to disconnect a volume from a server
func (c *Client) DetachVolume(ctx context.Context, identifier string) (*Volume, error) {
	return APIPost[Volume](
		ctx,
		c,
		path.Join(VolumeAPIPath, identifier, "detach"),
		nil,
	)
}

// ResizeVolume issues a request to change the size of a volume.
// The old size has to be specified as well as the new one.
func (c *Client) ResizeVolume(ctx context.Context, identifier string, newSize VolumeNewSize) (*Volume, error) {
	return APIPost[Volume](
		ctx,
		c,
		path.Join(VolumeAPIPath, identifier, "resize"),
		newSize,
	)
}

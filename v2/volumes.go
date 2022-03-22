package brightbox

import (
	"context"
	"path"

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
	ID               string  `json:"-"`
	Name             *string `json:"name,omitempty"`
	Serial           *string `json:"name,omitempty"`
	Size             *int    `json:"size,omitempty"`
	Image            *string `json:"image,omitempty"`
	Encrypted        *bool   `json:"encrypted,omitempty"`
	DeleteWithServer *bool   `json:"delete_with_server,omitempty"`
}

// VolumeAttachment is used in conjunction with AttachVolume and DetachVolume
type VolumeAttachment struct {
	Server string `json:"server"`
	Boot   bool   `json:"boot"`
}

// AttachVolume issues a request to attach the volume to a particular server and
// optionally mark it as the boot volume
func (c *Client) AttachVolume(ctx context.Context, identifier string, server VolumeAttachment) (*Volume, error) {
	return APIPost[Volume](
		ctx,
		c,
		path.Join(VolumeAPIPath, identifier, "attach"),
		server,
	)
}

// DetachVolume issues a request to disconnect a volume from a server
func (c *Client) DetachVolume(ctx context.Context, identifier string) error {
	return APIPostCommand(
		ctx,
		c,
		path.Join(VolumeAPIPath, identifier, "detach"),
	)
}

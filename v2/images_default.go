// Code generated by go generate; DO NOT EDIT.

package brightbox

import "context"
import "path"

const (
	// ImageAPIPath returns the relative URL path to the Image endpoint
	ImageAPIPath = "images"
)

// Images returns the collection view for Image
func (c *Client) Images(ctx context.Context) ([]Image, error) {
	return APIGetCollection[Image](ctx, c, ImageAPIPath)
}

// Image retrieves a detailed view of one resource
func (c *Client) Image(ctx context.Context, identifier string) (*Image, error) {
	return APIGet[Image](ctx, c, path.Join(ImageAPIPath, identifier))
}

// CreateImage creates a new resource from the supplied option map.
//
// It takes an instance of ImageOptions. Not all attributes can be
// specified at create time (such as ID, which is allocated for you).
func (c *Client) CreateImage(ctx context.Context, newImage *ImageOptions) (*Image, error) {
	return APIPost[Image](ctx, c, ImageAPIPath, newImage)
}

// UpdateImage updates an existing resources's attributes. Not all
// attributes can be changed (such as ID).
//
// It takes an instance of ImageOptions. Specify the resource you
// want to update using the ID field.
func (c *Client) UpdateImage(ctx context.Context, updateImage *ImageOptions) (*Image, error) {
	return APIPut[Image](ctx, c, path.Join(ImageAPIPath, updateImage.ID), updateImage)
}

// DestroyImage destroys an existing resource.
func (c *Client) DestroyImage(ctx context.Context, identifier string) error {
	return APIDelete(ctx, c, path.Join(ImageAPIPath, identifier))
}

// LockImage locks a resource against destroy requests
func (c *Client) LockImage(ctx context.Context, identifier string) error {
	return APIPutCommand(ctx, c, path.Join(ImageAPIPath, identifier, "lock_resource"))
}

// UnlockImage unlocks a resource, re-enabling destroy requests
func (c *Client) UnlockImage(ctx context.Context, identifier string) error {
	return APIPutCommand(ctx, c, path.Join(ImageAPIPath, identifier, "unlock_resource"))
}
package gobrightbox

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

type ImageOptions struct {
	Arch              string `json:"arch"`
	CompatibilityMode bool   `json:"compatibility_mode"`
	Description       string `json:"description"`
	MinRam            int    `json:"min_ram"`
	Name              string `json:"name"`
	Public            bool   `json:"public"`
	Username          string `json:"username"`
	Server            string `json:"server,omitempty"`
	Volume            string `json:"volume,omitempty"`
	HTTPURL           string `json:"http_url,omitempty"`
}

// Images retrieves a list of all images
func (c *Client) Images() ([]Image, error) {
	var images []Image
	_, err := c.MakeAPIRequest("GET", "/1.0/images", nil, &images)
	if err != nil {
		return nil, err
	}
	return images, err
}

// Image retrieves a detailed view of one image
func (c *Client) Image(identifier string) (*Image, error) {
	image := new(Image)
	_, err := c.MakeAPIRequest("GET", "/1.0/images/"+identifier, nil, image)
	if err != nil {
		return nil, err
	}
	return image, err
}

// DestroyImage issues a request to destroy the image
func (c *Client) DestroyImage(identifier string) error {
	_, err := c.MakeAPIRequest("DELETE", "/1.0/images/"+identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreateImage issues a request to create an image
func (c *Client) CreateImage(newImage *ImageOptions) (*Image, error) {
	image := new(Image)
	if _, err := c.MakeAPIRequest("POST", "/1.0/images", newImage, &image); err != nil {
		return nil, err
	}

	return image, nil
}

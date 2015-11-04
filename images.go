package brightbox

import (
	"time"
)

type Image struct {
	Resource
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
	AncestorId        string `json:"ancestor_id"`
	LicenseName       string `json:"license_name"`
}

func (c *Client) Images() ([]Image, error) {
	images := new([]Image)
	_, err := c.MakeApiRequest("GET", "/1.0/images", nil, images)
	if err != nil {
		return nil, err
	}
	return *images, err
}

func (c *Client) Image(identifier string) (*Image, error) {
	image := new(Image)
	_, err := c.MakeApiRequest("GET", "/1.0/images/"+identifier, nil, image)
	if err != nil {
		return nil, err
	}
	return image, err
}

func (c *Client) DestroyImage(identifier string) error {
	_, err := c.MakeApiRequest("DELETE", "/1.0/images/"+identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

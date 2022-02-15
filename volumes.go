package gobrightbox

// Volume represents a Brightbox Volume
// https://api.gb1.brightbox.com/1.0/#volume
type Volume struct {
	ID          string
	Name        string
	Status      string
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
	ID    string  `json:"-"`
	Size  *int    `json:"size,omitempty"`
	Image *string `json:"image,omitempty"`
}

// VolumeResizeOptions is used to change the size of a volume
type VolumeResizeOptions struct {
	From int
	To   int
}

// ResizeVolume changes the size of a volume
func (c *Client) ResizeVolume(identifier string, options *VolumeResizeOptions) error {
	_, err := c.MakeAPIRequest("POST", "/1.0/volumes/"+identifier+"/resize", options, nil)
	return err
}

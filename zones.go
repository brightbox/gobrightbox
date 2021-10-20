package gobrightbox

import (
	"fmt"
)

// Zone represents a Zone
// https://api.gb1.brightbox.com/1.0/#zone
type Zone struct {
	ID     string
	Handle string
}

// Zones retrieves a list of all Zones
func (c *Client) Zones() ([]Zone, error) {
	var zones []Zone
	_, err := c.MakeAPIRequest("GET", "/1.0/zones", nil, &zones)
	if err != nil {
		return nil, err
	}
	return zones, err
}

// Zone retrieves a detailed view of one Zone using an identifier
func (c *Client) Zone(identifier string) (*Zone, error) {
	zone := new(Zone)
	_, err := c.MakeAPIRequest("GET", "/1.0/zones/"+identifier, nil, zone)
	if err != nil {
		return nil, err
	}
	return zone, err
}

// ZoneByHandle retrieves a detailed view of one Zone using a handle
func (c *Client) ZoneByHandle(handle string) (*Zone, error) {
	zones, err := c.Zones()
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		if zone.Handle == handle {
			return &zone, nil
		}
	}
	return nil, fmt.Errorf("Zone with handle '%s' doesn't exist", handle)
}

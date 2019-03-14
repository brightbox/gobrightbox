package swauth

import "github.com/gophercloud/gophercloud"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "v1"
}

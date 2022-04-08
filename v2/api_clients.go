package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/permissionsgroup"
)

//go:generate ./generate_status_enum permissionsgroup full storage

// APIClient represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type APIClient struct {
	ResourceRef
	ID               string
	Name             string
	Description      string
	Secret           string
	PermissionsGroup permissionsgroup.Status `json:"permissions_group"`
	RevokedAt        *time.Time              `json:"revoked_at"`
	Account          *Account
}

// APIClientOptions is used to create and update api clients
type APIClientOptions struct {
	ID               string                  `json:"-"`
	Name             *string                 `json:"name,omitempty"`
	Description      *string                 `json:"description,omitempty"`
	PermissionsGroup permissionsgroup.Status `json:"permissions_group,omitempty"`
}

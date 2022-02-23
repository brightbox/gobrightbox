package brightbox

import (
	"time"

	"github.com/brightbox/gobrightbox/v2/status/collaboration"
)

//go:generate ./generate_status_enum collaboration pending accepted rejected cancelled ended

// Collaboration represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type Collaboration struct {
	ResourceRef
	ID         string
	Email      string
	Role       string
	RoleLabel  string `json:"role_label"`
	Status     collaboration.Status
	CreatedAt  *time.Time `json:"created_at"`
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	Account    *Account
	User       *User
	Inviter    *User
}

// CollaborationOptions is used to create and update api clients
type CollaborationOptions struct {
	ID    string  `json:"-"`
	Email *string `json:"email,omitempty"`
	Role  *string `json:"role,omitempty"`
}

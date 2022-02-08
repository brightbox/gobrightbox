package brightbox

import (
	"time"
)

// Collaboration represents an API client.
// https://api.gb1.brightbox.com/1.0/#api_client
type Collaboration struct {
	ResourceRef
	ID         string
	Email      string
	Role       string
	RoleLabel  string `json:"role_label"`
	Status     string
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

// APIPath returns the relative URL path to the collection endpoint
func (c Collaboration) APIPath() string {
	return "collaborations"
}

// FetchID returns the ID field from the object
func (c Collaboration) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c Collaboration) PostPath(from *CollaborationOptions) string {
	return c.APIPath()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c Collaboration) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c CollaborationOptions) OptionID() string {
	return c.ID
}

package brightbox

import (
	"time"
)

// Server represents a Cloud Server
// https://api.gb1.brightbox.com/1.0/#server
// DeletedAt is nil if the server has not yet been deleted
type Server struct {
	ResourceRef
	ID                string
	Name              string
	Status            string
	Hostname          string
	Fqdn              string
	UserData          string     `json:"user_data"`
	CreatedAt         *time.Time `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	StartedAt         *time.Time `json:"started_at"`
	Locked            bool       `json:"locked"`
	CompatibilityMode bool       `json:"compatibility_mode"`
	DiskEncrypted     bool       `json:"disk_encrypted"`
	Account           *Account
	Image             *Image
	Zone              *Zone
	ServerType        *ServerType   `json:"server_type"`
	CloudIPs          []CloudIP     `json:"cloud_ips"`
	ServerGroups      []ServerGroup `json:"server_groups"`
	Snapshots         []Server
	Interfaces        []Interface
	Volumes           []Volume
	ServerConsole
}

// ServerConsole is embedded into Server and contains the fields used in response
// to an ActivateConsoleForServer request.
type ServerConsole struct {
	ConsoleToken        string     `json:"console_token"`
	ConsoleURL          string     `json:"console_url"`
	ConsoleTokenExpires *time.Time `json:"console_token_expires"`
}

// ServerOptions is used in conjunction with CreateServer and UpdateServer to
// create and update servers.
type ServerOptions struct {
	ID                string   `json:"-"`
	Server            *string  `json:"image,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ServerType        *string  `json:"server_type,omitempty"`
	Zone              *string  `json:"zone,omitempty"`
	UserData          *string  `json:"user_data,omitempty"`
	ServerGroups      []string `json:"server_groups,omitempty"`
	CompatibilityMode *bool    `json:"compatibility_mode,omitempty"`
	DiskEncrypted     *bool    `json:"disk_encrypted,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c Server) APIPath() string {
	return "servers"
}

// FetchID returns the ID field from the object
func (c Server) FetchID() string {
	return c.ID
}

// PostPath returns the relative URL path to POST an object
func (c Server) PostPath(from *ServerOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c Server) PutPath(from *ServerOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c Server) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c ServerOptions) OptionID() string {
	return c.ID
}

// LockID returns the path to a lockable object
func (c Server) LockID() string {
	return c.APIPath() + "/" + c.FetchID()
}

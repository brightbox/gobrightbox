package brightbox

import (
	"context"
	"path"
	"time"

	"github.com/brightbox/gobrightbox/v2/status/server"
)

//go:generate ./generate_status_enum server creating active inactive deleting deleted failed unavailable

// Server represents a Cloud Server
// https://api.gb1.brightbox.com/1.0/#server
// DeletedAt is nil if the server has not yet been deleted
type Server struct {
	ResourceRef
	ServerConsole
	ID                      string
	Name                    string
	Status                  server.Status
	Hostname                string
	Fqdn                    string
	UserData                string     `json:"user_data"`
	CreatedAt               *time.Time `json:"created_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	StartedAt               *time.Time `json:"started_at"`
	SnapshotsSchedule       *string    `json:"snapshots_schedule"`
	SnapshotsScheduleNextAt *time.Time `json:"snapshots_schedule_next_at"`
	SnapshotsRetention      *string    `json:"snapshots_retention"`
	Locked                  bool       `json:"locked"`
	CompatibilityMode       bool       `json:"compatibility_mode"`
	DiskEncrypted           bool       `json:"disk_encrypted"`
	Account                 *Account
	Image                   *Image
	Zone                    *Zone
	ServerType              *ServerType   `json:"server_type"`
	CloudIPs                []CloudIP     `json:"cloud_ips"`
	ServerGroups            []ServerGroup `json:"server_groups"`
	Snapshots               []Server
	Interfaces              []Interface
	Volumes                 []Volume
}

// ServerConsole is embedded into Server and contains the fields used in response
// to an ActivateConsoleForServer request.
type ServerConsole struct {
	ConsoleToken        *string    `json:"console_token"`
	ConsoleURL          *string    `json:"console_url"`
	ConsoleTokenExpires *time.Time `json:"console_token_expires"`
}

// ServerOptions is used in conjunction with CreateServer and UpdateServer to
// create and update servers.
type ServerOptions struct {
	ID                string          `json:"-"`
	Image             *string         `json:"image,omitempty"`
	Name              *string         `json:"name,omitempty"`
	ServerType        *string         `json:"server_type,omitempty"`
	Zone              *string         `json:"zone,omitempty"`
	UserData          *string         `json:"user_data,omitempty"`
	ServerGroups      []string        `json:"server_groups,omitempty"`
	CompatibilityMode *bool           `json:"compatibility_mode,omitempty"`
	DiskEncrypted     *bool           `json:"disk_encrypted,omitempty"`
	Volumes           []VolumeOptions `json:"volumes,omitempty"`
}

// ActivateConsoleForServer issues a request to enable the graphical console for
// an existing server. The temporarily allocated ConsoleURL, ConsoleToken and
// ConsoleTokenExpires data are returned within an instance of Server.
func (c *Client) ActivateConsoleForServer(ctx context.Context, identifier string) (*Server, error) {
	return APIPost[Server](
		ctx,
		c,
		path.Join(ServerAPIPath, identifier, "activate_console"),
		nil,
	)
}

// ResizeServer issues a request to change the server type of a server
// changing the amount of cpu and ram it has.
func (c *Client) ResizeServer(ctx context.Context, identifier string, newTypeID string) error {
	return APIPostForm(
		ctx,
		c,
		path.Join(ServerAPIPath, identifier, "resize"),
		map[string]string{"new_type": newTypeID},
	)
}

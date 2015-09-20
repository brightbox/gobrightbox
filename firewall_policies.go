package brightbox

import (
	"time"
)

type FirewallPolicy struct {
	Resource
	Name        string
	Default     bool
	CreatedAt   *time.Time `json:"created_at"`
	Description *string
}

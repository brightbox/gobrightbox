package brightbox

import (
	"context"
)

type lockable interface {
	FetchID() string
	LockID() string
}

// LockResource locks a resource against destroy requests.
func LockResource(ctx context.Context, q *Client, resource lockable) error {
	return APIPutCommand(ctx, q, resource.LockID()+"/lock_resource")
}

// UnlockResource unlocks a resource, renabling destroy requests.
func UnlockResource(ctx context.Context, q *Client, resource lockable) error {
	return APIPutCommand(ctx, q, resource.LockID()+"/unlock_resource")
}

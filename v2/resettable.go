package brightbox

import (
	"context"
)

type resettable interface {
	ResetPasswordPath() string
}

// ResetPassword issues the reset command for the identified resource
// This is used to reset passwords for the requested resource type
// The returned resource is the only time the password is available. 
func ResetPassword[T resettable](ctx context.Context, q *Client, resource *T) (*T, error) {
	return APIPost[T](
		ctx,
		q,
		(*resource).ResetPasswordPath(),
		nil,
	)
}

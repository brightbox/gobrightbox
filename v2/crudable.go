package brightbox

import (
	"context"
)

type optionID interface{
	OptionID() string
}

type createable[I optionID] interface {
	FetchID() string
	PostPath(from *I) string
}

type updateable[I optionID] interface {
	FetchID() string
	PutPath(from *I) string
}

type destroyable interface {
	DestroyPath(from string) string
}

// Create creates a new resource from the supplied option map
//
// It takes an instance of Options. Not all attributes can be
// specified at create time (such as ID, which is allocated for you).
func Create[T createable[I], I optionID](ctx context.Context, q *Client, newOptions *I) (*T, error) {
	var zero T
	return APIPost[T](
		ctx,
		q,
		zero.PostPath(newOptions),
		newOptions,
	)
}

// Update updates an existing resources's attributes. Not all
// attributes can be changed (such as ID).
//
// Specify the resource you want to update using the ID field
// field.
func Update[T updateable[I], I optionID](ctx context.Context, q *Client, updateOptions *I) (*T, error) {
	var zero T
	return APIPut[T](
		ctx,
		q,
		zero.PutPath(updateOptions),
		updateOptions,
	)
}

// Destroy destroys an existing resource.
func Destroy[T destroyable](ctx context.Context, q *Client, identifier string) error {
	var zero T
	return APIDelete(
		ctx,
		q,
		zero.DestroyPath(identifier),
	)
}

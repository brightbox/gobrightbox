package brightbox

import (
	"context"
)

// Local interface type specifying the minimum API interface for a Brightbox resource
type queriable interface {
	FetchID() string
	APIPath() string
}

// All returns the result of making a collection call to the Brightbox API
// for the instantiated Brightbox resource. 
func All[T queriable](ctx context.Context, q *Client) ([]T, error) {
	var zero T
	collection, err := APIGet[[]T](
		ctx,
		q,
		zero.APIPath(),
	)
	if collection != nil {
		return *collection, err
	}
	return nil, err
}

// Instance retrieves a detailed view of one resource instance
func Instance[T queriable](ctx context.Context, q *Client, identifier string) (*T, error) {
	var zero T
	return APIGet[T](
		ctx,
		q,
		zero.APIPath() + "/" + identifier,
	)
}

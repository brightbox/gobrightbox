package brightbox

// Local interface type specifying the minimum API interface for a Brightbox resource
type queriable interface {
	FetchID() string
	APIPath() string
}

// All returns the result of making a collection call to the Brightbox API
// for the instantiated Brightbox resource. 
func All[T queriable](q *Client) ([]T, error) {
	var collection []T
	var zero T
	_, err := q.MakeAPIRequest(
		"GET",
		zero.APIPath(),
		nil,
		&collection,
	)
	return collection, err
}

// Instance retrieves a detailed view of one resource instance
func Instance[T queriable](q *Client, identifier string) (*T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"GET",
		resource.APIPath() + "/" + identifier,
		nil,
		&resource,
	)
	return &resource, err
}

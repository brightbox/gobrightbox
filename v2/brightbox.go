package brightbox

import (
	"context"
	"net/http"
	"net/url"
)

// Abstract Interface of any Brightbox oauth2 client generator
type oauth2 interface {
	Client(ctx context.Context) (*http.Client, error)
	APIURL() (*url.URL, error)
}

// Connect allocates and configures a Client for interacting with the API.
func Connect(ctx context.Context, config oauth2) (*Client, error) {
	baseURL, err := config.APIURL()
	if err != nil {
		return nil, err
	}
	httpClient, err := config.Client(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseURL: baseURL,
		client:  httpClient,
	}, nil
}

// Local interface type specifying the minimum API interface for a Brightbox resource
type queriable interface {
	APIPath() string
}

type optionID interface{
	FetchID() string
}

// All returns the result of making a collection call to the Brightbox API
// for the instantiated Brightbox resource. 
func All[T queriable](q *Client) ([]T, error) {
	var collection []T
	var resource T
	_, err := q.MakeAPIRequest(
		"GET",
		resource.APIPath(),
		nil,
		&collection,
	)
	return collection, err
}

// Instance retrieves a detailed view of one resource instance
func Instance[T queriable](q *Client, identifier string) (T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"GET",
		resource.APIPath() + "/" + identifier,
		nil,
		&resource,
	)
	return resource, err
}

type crudable[I optionID] interface {
	queriable
	Extract() I
}

// Create creates a new resource from the supplied option map
//
// It takes an instance of Options. Not all attributes can be
// specified at create time (such as ID, which is allocated for you).
func Create[T crudable[I], I optionID](q *Client, newOptions I) (T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"POST",
		resource.APIPath(),
		newOptions,
		&resource,
	)
	return resource, err
}

// Update updates an existing resources's attributes. Not all
// attributes can be changed (such as ID).
//
// Specify the resource you want to update using the ID field
// field.
func Update[T crudable[I], I optionID](q *Client, updateOptions I) (T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"PUT",
		resource.APIPath() + "/" + updateOptions.FetchID(),
		updateOptions,
		&resource,
	)
	return resource, err
}

// Destroy destroys an existing resource.
func Destroy[T queriable](q *Client, identifier string) error {
	var resource T
	_, err := q.MakeAPIRequest(
		"DELETE",
		resource.APIPath() + "/" + identifier,
		nil,
		nil,
	)
	return err
}

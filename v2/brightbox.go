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

// Client represents a connection to the Brightbox API. You should use NewConnect
// to allocate and configure Clients, and pass in either a
// clientcredentials or password configuration.
type Client struct {
	UserAgent string
	baseURL   *url.URL
	client    *http.Client
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
type resource interface {
}

// Querier is a generic facilitator type that can be instantied with any
// Brightbox Resource type
type Querier[T resource] struct {
	client *Client
}

// NewQuerier returns a client that is primed to interact with a
// particular Brightbox Resource.
func NewQuerier[T resource](c *Client) *Querier[T] {
	return &Querier[T]{
		client: c,
	}
}

// All returns the result of making a collection call to the Brightbox API
// for the instantiated Brightbox resource. 
func (q *Querier[T]) All() ([]T, error) {
	// implementation
}

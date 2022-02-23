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

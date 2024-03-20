package brightbox

import (
	"context"
	"net/http/httptest"
)

func SetupConnection(handler *APIMock) (*httptest.Server, *Client, error) {
	ts := httptest.NewServer(handler)

	// Setup Mock Config
	conf := &MockAuth{
		url: ts.URL,
	}

	// Underlying network connection context.
	ctx := context.Background()

	// Setup connection to API
	client, err := Connect(ctx, conf)
	client.DisallowUnknownFields()
	return ts, client, err
}

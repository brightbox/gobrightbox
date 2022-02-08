package brightbox

import (
	"context"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
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

func TestResetSecret(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients/cli-dsse2/reset_secret",
			ExpectBody:   "",
			GiveBody:     readJSON("api_client"),
		},
	)
	assert.NilError(t, err)
	defer ts.Close()

	nc, err := ResetSecret(context.Background(), client, &APIClient{ID: "cli-dsse2"})
	assert.NilError(t, err)
	assert.Assert(t, nc != nil)
}

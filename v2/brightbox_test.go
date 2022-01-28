package brightbox

import (
	"context"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
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
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	nc, err := ResetSecret(client, &APIClient{ID: "cli-dsse2"})
	assert.Assert(t, is.Nil(err), "ResetSecret() returned an error")
	assert.Assert(t, nc != nil, "ResetSecret() returned nil")
}

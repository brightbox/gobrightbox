package brightbox

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func testAll[I any](
	t *testing.T,
	allInstances func(c *Client, ctx context.Context) ([]I, error),
	typeName string,
	apiPath string,
	instanceRef string,
) *I {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/" + apiPath,
			ExpectBody:   "",
			GiveBody:     readJSON(apiPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	collection, err := allInstances(client, context.Background())
	assert.Assert(t, is.Nil(err), "All["+typeName+"] returned an error")
	assert.Assert(t, collection != nil, "All["+typeName+"] returned nil")
	assert.Equal(t, 1, len(collection), "wrong number of "+instanceRef+" returned")
	return &collection[0]
}

func testInstance[I any](
	t *testing.T,
	fetchInstance func(*Client, context.Context, string) (*I, error),
	typeName string,
	expectPath string,
	jsonPath string,
	instanceID string,
) *I {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/" + expectPath,
			ExpectBody:   "",
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	instance, err := fetchInstance(client, context.Background(), instanceID)
	assert.Assert(t, is.Nil(err), "Instance["+typeName+"] returned an error")
	assert.Assert(t, instance != nil, "Instance["+typeName+"] returned nil")
	return instance
}

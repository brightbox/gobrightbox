package brightbox

import (
	"context"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testByHandle[I handleable](
	t *testing.T,
	typeName string,
	apiPath string,
	instanceID string,
	handle string,
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

	instance, err := ByHandle[I](context.Background(), client, handle)
	assert.Assert(t, is.Nil(err), "Instance[" + typeName + "] returned an error")
	assert.Assert(t, instance != nil, "Instance[" + typeName + "] returned nil")
	assert.Equal(t, (*instance).FetchID(), instanceID)
	return instance
}

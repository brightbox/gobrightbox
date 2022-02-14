package brightbox

import (
	"context"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testResetPassword[O resettable](
	t *testing.T,
	typeName string,
	apiPath string,
	jsonPath string,
	instance *O,
	resetEndpoint string,
) *O {
	instanceID := (*instance).FetchID()
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/" + apiPath + "/" + instanceID + "/" + resetEndpoint,
			ExpectBody:   "",
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	newRes, err := ResetPassword[O](context.Background(), client, instance)
	assert.Assert(t, is.Nil(err), "Reset[" + typeName + "] returned an error")
	assert.Assert(t, newRes != nil, "Reset[" + typeName + "] returned nil")
	return newRes
}

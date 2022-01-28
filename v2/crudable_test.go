package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testCreate[O crudable[I], I optionID](
	t *testing.T,
	typeName string,
	apiPath string,
	jsonPath string,
	instanceID string,
	newOptions *I,
	expectedBody string,
) *O {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/" + apiPath,
			ExpectBody:   expectedBody,
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	instance, err := Create[O](client, newOptions)
	assert.Assert(t, is.Nil(err), "Create[" + typeName + "] returned an error")
	assert.Assert(t, instance != nil, "Create[" + typeName + "] returned nil")
	assert.Equal(t, instanceID, (*(*instance).Extract()).FetchID())
	return instance
}


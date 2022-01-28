package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testAll[I queriable](
	t *testing.T,
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

	collection, err := All[I](client)
	assert.Assert(t, is.Nil(err), "All["+ typeName + "] returned an error")
	assert.Assert(t, collection != nil, "All[" + typeName + "] returned nil")
	assert.Equal(t, 1, len(collection), "wrong number of "+ instanceRef + "s returned")
	return &collection[0]
}

func testInstance[I queriable](
	t *testing.T,
	typeName string,
	apiPath string,
	jsonPath string,
	instanceID string,
) *I {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/" + apiPath + "/" + instanceID,
			ExpectBody:   "",
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	instance, err := Instance[I](client, instanceID)
	assert.Assert(t, is.Nil(err), "Instance[" + typeName + "] returned an error")
	assert.Assert(t, instance != nil, "Instance[" + typeName + "] returned nil")
	return instance
}

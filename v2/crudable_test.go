package brightbox

import (
	"context"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testCreate[O createable[I], I optionID](
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
	instance, err := Create[O](context.Background(), client, newOptions)
	assert.Assert(t, is.Nil(err), "Create[" + typeName + "] returned an error")
	assert.Assert(t, instance != nil, "Create[" + typeName + "] returned nil")
	assert.Equal(t, instanceID, (*instance).FetchID())
	return instance
}

func testUpdate[O updateable[I], I optionID](
	t *testing.T,
	typeName string,
	apiPath string,
	jsonPath string,
	instanceID string,
	updatedOptions *I,
	expectedBody string,
) *O {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/" + apiPath + "/" + instanceID,
			ExpectBody:   expectedBody,
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	instance, err := Update[O](context.Background(), client, updatedOptions)
	assert.Assert(t, is.Nil(err), "Update[" + typeName + "] returned an error")
	assert.Assert(t, instance != nil, "Update[" + typeName + "] returned nil")
	assert.Equal(t, instanceID, (*instance).FetchID())
	return instance
}

func testDestroy[O destroyable](
	t *testing.T,
	typeName string,
	apiPath string,
	instanceID string,
) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "DELETE",
			ExpectURL:    "/1.0/" + apiPath + "/" + instanceID,
			ExpectBody:   "",
			GiveBody:     "",
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	err = Destroy[O](context.Background(), client, instanceID)
	assert.Assert(t, is.Nil(err), "Destroy[" + typeName + "] returned an error")
}

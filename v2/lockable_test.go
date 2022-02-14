package brightbox

import (
	"context"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func testLock[I lockable](
	t *testing.T,
	typeName string,
	apiPath string,
	instance *I,
	lock_direction string,
	lock_function func (
		context.Context,
		*Client,
		lockable,
	) error,
) {
	instanceID := (*instance).FetchID()
	assert.Assert(t, instanceID != "")
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/" + apiPath + "/" + instanceID + "/" + lock_direction,
			ExpectBody:   ``,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = lock_function(
		context.Background(),
		client,
		*instance,
	)
	assert.Assert(t, is.Nil(err), typeName + " " +lock_direction + " returned an error")
}

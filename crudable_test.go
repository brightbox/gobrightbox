package brightbox

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func testModify[O, I any](
	t *testing.T,
	modify func(*Client, context.Context, I) (*O, error),
	newOptions I,
	jsonPath string,
	verb string,
	expectedPath string,
	expectedBody string,
) *O {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: verb,
			ExpectURL:    "/1.0/" + expectedPath,
			ExpectBody:   expectedBody,
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	instance, err := modify(client, context.Background(), newOptions)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, instance != nil)
	return instance
}

func testLink[O, I any](
	t *testing.T,
	modify func(*Client, context.Context, string, I) (*O, error),
	from string,
	to I,
	jsonPath string,
	verb string,
	expectedPath string,
	expectedBody string,
) *O {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: verb,
			ExpectURL:    "/1.0/" + expectedPath,
			ExpectBody:   expectedBody,
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	instance, err := modify(client, context.Background(), from, to)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, instance != nil)
	return instance
}

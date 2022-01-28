package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestServerTypes(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/server_types",
			ExpectBody:   "",
			GiveBody:     readJSON("server_types"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	p, err := All[ServerType](client)
	assert.Assert(t, is.Nil(err), "All[ServerType]() returned an error")
	assert.Assert(t, p != nil, "All[ServerType]() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of server_types returned")
	ac := p[0]
	assert.Equal(t, "typ-zx45f", ac.ID, "server_type id incorrect")
}

func TestServerType(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/server_types/typ-zx45f",
			ExpectBody:   "",
			GiveBody:     readJSON("server_type"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	ac, err := Instance[ServerType](client, "typ-zx45f")
	assert.Assert(t, is.Nil(err), "Instance[ServerType] returned an error")
	assert.Assert(t, ac != nil, "Instance[ServerType] returned nil")
	assert.Equal(t, "typ-zx45f", ac.ID, "server_type id incorrect")
	assert.Equal(t, "Small", ac.Name, "server_type name incorrect")
}

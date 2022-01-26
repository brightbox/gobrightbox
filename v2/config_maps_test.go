package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestConfigMaps(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/config_maps",
			ExpectBody:   "",
			GiveBody:     readJSON("config_maps"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	p, err := All[ConfigMap](client)
	assert.Assert(t, is.Nil(err), "All[ConfigMaps] returned an error")
	assert.Assert(t, p != nil, "All[ConfigMaps] returned nil")
	assert.Equal(t, 1, len(p), "wrong number of config map returned")
	ac := p[0]
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID), "config map id incorrect")
}

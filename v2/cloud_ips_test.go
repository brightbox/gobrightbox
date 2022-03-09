package brightbox

import (
	"context"
	"path"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMapCloudIP(t *testing.T) {
	command := (*Client).MapCloudIP
	instanceID := "cip-k4a25"
	targetID := "lba-12345"
	verb := "POST"
	expectedPath := path.Join("cloud_ips", "cip-k4a25", "map")
	expectedBody := `{"destination":"lba-12345"}`

	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: verb,
			ExpectURL:    "/1.0/" + expectedPath,
			ExpectBody:   expectedBody,
			GiveBody:     "",
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	err = command(client, context.Background(), instanceID, targetID)
	assert.Assert(t, is.Nil(err))
}

func TestUnMapCloudIP(t *testing.T) {
	testCommand(
		t,
		(*Client).UnMapCloudIP,
		"cip-k4a25",
		"POST",
		path.Join("cloud_ips", "cip-k4a25", "unmap"),
	)
}

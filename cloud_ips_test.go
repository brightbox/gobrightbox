package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMapCloudIP(t *testing.T) {
	instance := testLink(
		t,
		(*Client).MapCloudIP,
		"cip-k4a25",
		CloudIPAttachment{"lba-12345"},
		"cloud_ip",
		"POST",
		path.Join("cloud_ips", "cip-k4a25", "map"),
		`{"destination":"lba-12345"}`,
	)
	assert.Equal(t, instance.ID, "cip-k4a25")
}

func TestUnMapCloudIP(t *testing.T) {
	instance := testModify(
		t,
		(*Client).UnMapCloudIP,
		"cip-k4a25",
		"cloud_ip",
		"POST",
		path.Join("cloud_ips", "cip-k4a25", "unmap"),
		"",
	)
	assert.Equal(t, instance.ID, "cip-k4a25")
}

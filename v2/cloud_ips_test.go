package brightbox

import (
	"path"
	"testing"
)

func TestMapCloudIP(t *testing.T) {
	testForm[string](
		t,
		(*Client).MapCloudIP,
		"cip-k4a25",
		"lba-12345",
		"POST",
		path.Join("cloud_ips", "cip-k4a25", "map"),
		`{"destination":"lba-12345"}`,
	)
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

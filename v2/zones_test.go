package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestZones(t *testing.T) {
	instance := testAll[Zone](
		t,
		"Zone",
		"zones",
		"zone",
	)
	assert.Equal(t, instance.ID, "zon-328ds")
}

func TestZone(t *testing.T) {
	instance := testInstance[Zone](
		t,
		"Zone",
		"zones",
		"zone",
		"zon-328ds",
	)
	assert.Equal(t, instance.Handle, "gb1")
}

func TestZoneByHandle(t *testing.T) {
	_ = testByHandle[Zone](
		t,
		"Zone",
		"zones",
		"zon-328ds",
		"gb1",
	)
}

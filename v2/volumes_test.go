package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestVolumes(t *testing.T) {
	instance := testAll[Volume](
		t,
		"Volume",
		"volumes",
		"volume",
	)
	assert.Equal(t, instance.ID, "vol-po5we")
}

func TestVolume(t *testing.T) {
	instance := testInstance[Volume](
		t,
		"Volume",
		"volumes",
		"volume",
		"vol-po5we",
	)
	assert.Equal(t, instance.StorageType, "network")
}

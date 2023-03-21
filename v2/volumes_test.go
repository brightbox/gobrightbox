package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAttachVolume(t *testing.T) {
	instance := testLink(
		t,
		(*Client).AttachVolume,
		"vol-po5we",
		VolumeAttachment{"srv-lv426", false},
		"volume",
		"POST",
		path.Join("volumes", "vol-po5we", "attach"),
		`{"server":"srv-lv426","boot":false}`,
	)
	assert.Equal(t, instance.ID, "vol-po5we")
}

func TestDetachVolume(t *testing.T) {
	instance := testModify(
		t,
		(*Client).DetachVolume,
		"vol-po5we",
		"volume",
		"POST",
		path.Join("volumes", "vol-po5we", "detach"),
		"",
	)
	assert.Equal(t, instance.ID, "vol-po5we")
}

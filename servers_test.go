package brightbox

import (
	"testing"

	"github.com/brightbox/gobrightbox/v2/enums/serverstatus"
	"gotest.tools/v3/assert"
)

func TestCreateServerWithNetworkDisk(t *testing.T) {
	name := "myserver"
	image := "img-linux"
	var size uint = 12345
	newAC := ServerOptions{
		Name: &name,
		Volumes: []VolumeEntry{
			VolumeEntry{
				Size:  size,
				Image: image,
			},
		},
	}
	instance := testModify(
		t,
		(*Client).CreateServer,
		newAC,
		"server",
		"POST",
		"servers",
		`{"name":"myserver","volumes":[{"size":12345,"image":"img-linux"}]}`,
	)
	assert.Equal(t, instance.ID, "srv-lv426")
	assert.Assert(t, instance.Status == serverstatus.Active)
}

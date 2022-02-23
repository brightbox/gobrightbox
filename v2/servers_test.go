package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestCreateServerWithNetworkDisk(t *testing.T) {
	name := "myserver"
	image := "img-linux"
	size := 12345
	newAC := ServerOptions{
		Name: &name,
		Volumes: []VolumeOptions{
			VolumeOptions{
				Image: &image,
				Size:  &size,
			},
		},
	}
	instance := testModify[Server, ServerOptions](
		t,
		(*Client).CreateServer,
		&newAC,
		"server",
		"POST",
		"servers",
		`{"name":"myserver","volumes":[{"size":12345,"image":"img-linux"}]}`,
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

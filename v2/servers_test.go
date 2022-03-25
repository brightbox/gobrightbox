package brightbox

import (
	"path"
	"testing"

	"github.com/brightbox/gobrightbox/v2/status/server"
	"gotest.tools/v3/assert"
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
	instance := testModify[Server, *ServerOptions](
		t,
		(*Client).CreateServer,
		&newAC,
		"server",
		"POST",
		"servers",
		`{"name":"myserver","volumes":[{"size":12345,"image":"img-linux"}]}`,
	)
	assert.Equal(t, instance.ID, "srv-lv426")
	assert.Assert(t, instance.Status == server.Active)
}

func TestActivateConsoleForServer(t *testing.T) {
	instance := testModify[Server, string](
		t,
		(*Client).ActivateConsoleForServer,
		"srv-lv426",
		"server",
		"POST",
		path.Join("servers", "srv-lv426", "activate_console"),
		"",
	)
	assert.Equal(t, instance.ID, "srv-lv426")

}

func TestResizeServer(t *testing.T) {
	instance := testLink[Server, ServerNewSize](
		t,
		(*Client).ResizeServer,
		"srv-lv426",
		ServerNewSize{"typ-12345"},
		"server",
		"POST",
		path.Join("servers", "srv-lv426", "resize"),
		`{"new_type":"typ-12345"}`,
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestServers(t *testing.T) {
	instance := testAll[Server](
		t,
		(*Client).Servers,
		"Server",
		"servers",
		"Servers",
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

func TestServer(t *testing.T) {
	instance := testInstance[Server](
		t,
		(*Client).Server,
		"Server",
		path.Join("servers", "srv-lv426"),
		"server",
		"srv-lv426",
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

func TestCreateServer(t *testing.T) {
	newResource := ServerOptions{}
	instance := testModify[Server, *ServerOptions](
		t,
		(*Client).CreateServer,
		&newResource,
		"server",
		"POST",
		path.Join("servers"),
		"{}",
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

func TestUpdateServer(t *testing.T) {
	updatedResource := ServerOptions{ID: "srv-lv426"}
	instance := testModify[Server, *ServerOptions](
		t,
		(*Client).UpdateServer,
		&updatedResource,
		"server",
		"PUT",
		path.Join("servers", updatedResource.ID),
		"{}",
	)
	assert.Equal(t, instance.ID, updatedResource.ID)
}

func TestDestroyServer(t *testing.T) {
	testCommand(
		t,
		(*Client).DestroyServer,
		"srv-lv426",
		"DELETE",
		path.Join("servers", "srv-lv426"),
	)
}

func TestLockServer(t *testing.T) {
	testCommand(
		t,
		(*Client).LockServer,
		"srv-lv426",
		"PUT",
		path.Join("servers", "srv-lv426", "lock_resource"),
	)
}

func TestUnlockServer(t *testing.T) {
	testCommand(
		t,
		(*Client).UnlockServer,
		"srv-lv426",
		"PUT",
		path.Join("servers", "srv-lv426", "unlock_resource"),
	)
}
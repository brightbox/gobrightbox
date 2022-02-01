package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestServers(t *testing.T) {
	instance := testAll[Server](
		t,
		"Server",
		"servers",
		"server",
	)
	assert.Equal(t, instance.ID, "srv-lv426")
}

func TestServer(t *testing.T) {
	instance := testInstance[Server](
		t,
		"Server",
		"servers",
		"server",
		"srv-lv426",
	)
	assert.Equal(t, instance.Status, "active")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.Image.ID, "img-3ikco")
	assert.Equal(t, instance.ServerType.ID, "typ-zx45f")
	assert.Equal(t, instance.Zone.ID, "zon-328ds")
}

func TestCreateServer(t *testing.T) {
	newAC := ServerOptions{}
	_ = testCreate[Server](
		t,
		"Server",
		"servers",
		"server",
		"srv-lv426",
		&newAC,
		"{}",
	)
}

func TestUpdateServer(t *testing.T) {
	name := "dev client"
	uac := ServerOptions{ID: "srv-lv426", Name: &name}
	_ = testUpdate[Server](
		t,
		"Server",
		"servers",
		"server",
		"srv-lv426",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyServer(t *testing.T) {
	testDestroy[Server](
		t,
		"Server",
		"servers",
		"srv-lv426",
	)
}

func TestLockServer(t *testing.T) {
	testLock[Server](
		t,
		"Server",
		"servers",
		&Server{ID: "srv-lv426"},
		"srv-lv426",
		"lock_resource",
		LockResource,
	)
}

func TestUnlockServer(t *testing.T) {
	testLock[Server](
		t,
		"Server",
		"servers",
		&Server{ID: "srv-lv426"},
		"srv-lv426",
		"unlock_resource",
		UnlockResource,
	)
}

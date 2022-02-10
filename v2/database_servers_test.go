package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestDatabaseServers(t *testing.T) {
	instance := testAll[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"database_server",
	)
	assert.Equal(t, instance.ID, "dbs-123ab")
}

func TestDatabaseServer(t *testing.T) {
	instance := testInstance[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"database_server",
		"dbs-123ab",
	)
	assert.Equal(t, instance.Status, "active")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.DatabaseServerType.ID, "dbt-12345")
	assert.Equal(t, instance.Zone.ID, "zon-328ds")
}

func TestCreateDatabaseServer(t *testing.T) {
	newAC := DatabaseServerOptions{}
	_ = testCreate[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"database_server",
		"dbs-123ab",
		&newAC,
		"{}",
	)
}

func TestUpdateDatabaseServer(t *testing.T) {
	name := "dev client"
	uac := DatabaseServerOptions{ID: "dbs-123ab", Name: &name}
	_ = testUpdate[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"database_server",
		"dbs-123ab",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyDatabaseServer(t *testing.T) {
	testDestroy[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"dbs-123ab",
	)
}

func TestLockDatabaseServer(t *testing.T) {
	testLock[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		&DatabaseServer{ID: "dbs-123ab"},
		"dbs-123ab",
		"lock_resource",
		LockResource,
	)
}

func TestUnlockDatabaseServer(t *testing.T) {
	testLock[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		&DatabaseServer{ID: "dbs-123ab"},
		"dbs-123ab",
		"unlock_resource",
		UnlockResource,
	)
}

func TestDatabaseServerResetPassword(t *testing.T) {
	instance := testResetPassword[DatabaseServer](
		t,
		"DatabaseServer",
		"database_servers",
		"database_server",
		&DatabaseServer{ID: "dbs-123ab"},
		"dbs-123ab",
		"reset_password",
	)
	assert.Equal(t, instance.ID, "dbs-123ab")
	assert.Equal(t, instance.Status, "active")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.DatabaseServerType.ID, "dbt-12345")
	assert.Equal(t, instance.Zone.ID, "zon-328ds")
}

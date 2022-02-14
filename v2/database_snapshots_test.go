package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestDatabaseSnapshots(t *testing.T) {
	instance := testAll[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		"database_snapshot",
	)
	assert.Equal(t, instance.ID, "dbi-12345")
}

func TestDatabaseSnapshot(t *testing.T) {
	instance := testInstance[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		"database_snapshot",
		"dbi-12345",
	)
	assert.Equal(t, instance.Status, "available")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.DatabaseEngine, "mysql")
}

func TestUpdateDatabaseSnapshot(t *testing.T) {
	name := "dev client"
	uac := DatabaseSnapshotOptions{ID: "dbi-12345", Name: &name}
	_ = testUpdate[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		"database_snapshot",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyDatabaseSnapshot(t *testing.T) {
	testDestroy[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		"dbi-12345",
	)
}

func TestLockDatabaseSnapshot(t *testing.T) {
	testLock[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		&DatabaseSnapshot{ID: "dbi-12345"},
		"lock_resource",
		LockResource,
	)
}

func TestUnlockDatabaseSnapshot(t *testing.T) {
	testLock[DatabaseSnapshot](
		t,
		"DatabaseSnapshot",
		"database_snapshots",
		&DatabaseSnapshot{ID: "dbi-12345"},
		"unlock_resource",
		UnlockResource,
	)
}

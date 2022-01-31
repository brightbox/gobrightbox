package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestAPIClients(t *testing.T) {
	instance := testAll[APIClient](
		t,
		"APIClient",
		"api_clients",
		"api client",
	)
	assert.Equal(t, instance.ID, "cli-dsse2")
}

func TestAPIClient(t *testing.T) {
	instance := testInstance[APIClient](
		t,
		"APIClient",
		"api_clients",
		"api_client",
		"cli-dsse2",
	)
	assert.Equal(t, instance.Name, "dev client")
}

func TestCreateAPIClient(t *testing.T) {
	newAC := APIClientOptions{}
	_ = testCreate[APIClient](
		t,
		"APIClient",
		"api_clients",
		"api_client",
		"cli-dsse2",
		&newAC,
		"{}",
	)
}

func TestCreateAPIClientWithPermissionsGroup(t *testing.T) {
	pg := "full"
	newAC := APIClientOptions{PermissionsGroup: &pg}
	instance := testCreate[APIClient](
		t,
		"APIClient",
		"api_clients",
		"api_client",
		"cli-dsse2",
		&newAC,
		`{"permissions_group":"full"}`,
	)
	assert.Equal(t, instance.PermissionsGroup, pg)
}

func TestUpdateAPIClient(t *testing.T) {
	name := "dev client"
	uac := APIClientOptions{ID: "cli-dsse2", Name: &name}
	_ = testUpdate[APIClient](
		t,
		"APIClient",
		"api_clients",
		"api_client",
		"cli-dsse2",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyAPIClient(t *testing.T) {
	testDestroy[APIClient](
		t,
		"APIClient",
		"api_clients",
		"cli-dsse2",
	)
}

func TestLockAPIClient(t *testing.T) {
	testLock[APIClient](
		t,
		"APIClient",
		"api_clients",
		&APIClient{ID: "cli-dsse2"},
		"cli-dsse2",
		"lock_resource",
		LockResource,
	)
}

func TestUnlockAPIClient(t *testing.T) {
	testLock[APIClient](
		t,
		"APIClient",
		"api_clients",
		&APIClient{ID: "cli-dsse2"},
		"cli-dsse2",
		"unlock_resource",
		UnlockResource,
	)
}

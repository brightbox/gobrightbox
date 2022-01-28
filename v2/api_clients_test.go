package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
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
	assert.Equal(t, instance.ID, "cli-dsse2")
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
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/api_clients/cli-dsse2/lock_resource",
			ExpectBody:   ``,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = LockResource(client, &APIClient{ID: "cli-dsse2"})
	assert.Assert(t, is.Nil(err), "LockAPIClient() returned an error")
}

func TestUnlockAPIClient(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/api_clients/cli-dsse2/unlock_resource",
			ExpectBody:   ``,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = UnLockResource(client, &APIClient{ID: "cli-dsse2"})
	assert.Assert(t, is.Nil(err), "LockAPIClient() returned an error")
}

package brightbox

import (
	"context"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func SetupConnection(handler *APIMock) (*httptest.Server, *Client, error) {
	ts := httptest.NewServer(handler)

	// Setup Mock Config
	conf := &MockAuth{
		url: ts.URL,
	}

	// Underlying network connection context.
	ctx := context.Background()

	// Setup connection to API
	client, err := Connect(ctx, conf)
	return ts, client, err
}

func TestAPIClients(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   "",
			GiveBody:     readJSON("api_clients"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	p, err := All[APIClient](client)
	assert.Assert(t, is.Nil(err), "All[APIClient]() returned an error")
	assert.Assert(t, p != nil, "All[APIClient]() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of api clients returned")
	ac := p[0]
	assert.Equal(t, "cli-dsse2", ac.ID, "api client id incorrect")
}

func TestAPIClient(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/api_clients/cli-dsse2",
			ExpectBody:   "",
			GiveBody:     readJSON("api_client"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	ac, err := Instance[APIClient](client, "cli-dsse2")
	assert.Assert(t, is.Nil(err), "Instance[APIClient] returned an error")
	assert.Assert(t, ac != nil, "Instance[APIClient] returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID, "api client id incorrect")
	assert.Equal(t, "dev client", ac.Name, "api client name incorrect")
}

func TestCreateAPIClient(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   "{}",
			GiveBody:     readJSON("api_client"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	newAC := APIClientOptions{}
	ac, err := Create[APIClient](client, &newAC)
	assert.Assert(t, is.Nil(err), "Create[APIClient] returned an error")
	assert.Assert(t, ac != nil, "Create[APIClient] returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
}

func TestCreateAPIClientWithPermissionsGroup(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   `{"permissions_group":"full"}`,
			GiveBody:     readJSON("api_client"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	pg := "full"
	newAC := APIClientOptions{PermissionsGroup: &pg}
	ac, err := Create[APIClient](client, &newAC)
	assert.Assert(t, is.Nil(err), "CreateAPIClient() returned an error")
	assert.Assert(t, ac != nil, "CreateAPIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
	assert.Equal(t, pg, ac.PermissionsGroup)
}

func TestUpdateAPIClient(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/api_clients/cli-dsse2",
			ExpectBody:   `{"name":"dev client"}`,
			GiveBody:     readJSON("api_client"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	name := "dev client"
	uac := APIClientOptions{ID: "cli-dsse2", Name: &name}
	ac, err := Update[APIClient](client, &uac)
	assert.Assert(t, is.Nil(err), "UpdateAPIClient() returned an error")
	assert.Assert(t, ac != nil, "UpdateAPIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
}

func TestDestroyAPIClient(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "DELETE",
			ExpectURL:    "/1.0/api_clients/cli-dsse2",
			ExpectBody:   "",
			GiveBody:     "",
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = Destroy[APIClient](client, "cli-dsse2")
	assert.Assert(t, is.Nil(err), "DestroyAPIClient() returned an error")
}

func TestResetSecret(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients/cli-dsse2/reset_secret",
			ExpectBody:   "",
			GiveBody:     readJSON("api_client"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	nc, err := ResetSecret(client, &APIClient{ID: "cli-dsse2"})
	assert.Assert(t, is.Nil(err), "ResetSecret() returned an error")
	assert.Assert(t, nc != nil, "ResetSecret() returned nil")
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

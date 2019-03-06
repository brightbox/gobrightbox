package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestApiClients(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   "",
		GiveBody:     readJSON("api_clients"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	p, err := client.ApiClients()
	require.Nil(t, err, "ApiClients() returned an error")
	require.NotNil(t, p, "ApiClients() returned nil")
	require.Equal(t, 1, len(p), "wrong number of api client returned")
	ac := p[0]
	assert.Equal(t, "cli-dsse2", ac.Id, "api client id incorrect")
}

func TestApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   "",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	ac, err := client.ApiClient("cli-dsse2")
	require.Nil(t, err, "ApiClient() returned an error")
	require.NotNil(t, ac, "ApiClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.Id, "api client id incorrect")
	assert.Equal(t, "dev client", ac.Name, "api client name incorrect")
}

func TestCreateApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   "{}",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	newAC := brightbox.ApiClientOptions{}
	ac, err := client.CreateApiClient(&newAC)
	require.Nil(t, err, "CreateApiClient() returned an error")
	require.NotNil(t, ac, "CreateApiClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.Id)
}

func TestCreateApiClientWithPermissionsGroup(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   `{"permissions_group":"full"}`,
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	pg := "full"
	newAC := brightbox.ApiClientOptions{PermissionsGroup: &pg}
	ac, err := client.CreateApiClient(&newAC)
	require.Nil(t, err, "CreateApiClient() returned an error")
	require.NotNil(t, ac, "CreateApiClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.Id)
	assert.Equal(t, pg, ac.PermissionsGroup)
}

func TestUpdateApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   `{"name":"dev client"}`,
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	name := "dev client"
	uac := brightbox.ApiClientOptions{Id: "cli-dsse2", Name: &name}
	ac, err := client.UpdateApiClient(&uac)
	require.Nil(t, err, "UpdateApiClient() returned an error")
	require.NotNil(t, ac, "UpdateApiClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.Id)
}

func TestDestroyApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.DestroyApiClient("cli-dsse2")
	require.Nil(t, err, "DestroyApiClient() returned an error")
}

func TestResetSecretForApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients/cli-dsse2/reset_secret",
		ExpectBody:   "",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	ac, err := client.ResetSecretForApiClient("cli-dsse2")
	require.Nil(t, err, "ResetPasswordForApiClient() returned an error")
	require.NotNil(t, ac, "ResetPasswordForApiClient() returned nil")
}

func TestLockApiClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/api_clients/cli-dsse2/lock_resource",
		ExpectBody:   ``,
		GiveBody:     ``,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.LockResource(brightbox.ApiClient{Id: "cli-dsse2"})
	require.Nil(t, err, "LockApiClient() returned an error")
}

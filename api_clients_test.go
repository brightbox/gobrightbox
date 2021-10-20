package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIClients(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   "",
		GiveBody:     readJSON("api_clients"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	p, err := client.APIClients()
	require.Nil(t, err, "APIClients() returned an error")
	require.NotNil(t, p, "APIClients() returned nil")
	require.Equal(t, 1, len(p), "wrong number of api client returned")
	ac := p[0]
	assert.Equal(t, "cli-dsse2", ac.ID, "api client id incorrect")
}

func TestAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   "",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	ac, err := client.APIClient("cli-dsse2")
	require.Nil(t, err, "APIClient() returned an error")
	require.NotNil(t, ac, "APIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID, "api client id incorrect")
	assert.Equal(t, "dev client", ac.Name, "api client name incorrect")
}

func TestCreateAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   "{}",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	newAC := gobrightbox.APIClientOptions{}
	ac, err := client.CreateAPIClient(&newAC)
	require.Nil(t, err, "CreateAPIClient() returned an error")
	require.NotNil(t, ac, "CreateAPIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
}

func TestCreateAPIClientWithPermissionsGroup(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   `{"permissions_group":"full"}`,
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	pg := "full"
	newAC := gobrightbox.APIClientOptions{PermissionsGroup: &pg}
	ac, err := client.CreateAPIClient(&newAC)
	require.Nil(t, err, "CreateAPIClient() returned an error")
	require.NotNil(t, ac, "CreateAPIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
	assert.Equal(t, pg, ac.PermissionsGroup)
}

func TestUpdateAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   `{"name":"dev client"}`,
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	name := "dev client"
	uac := gobrightbox.APIClientOptions{ID: "cli-dsse2", Name: &name}
	ac, err := client.UpdateAPIClient(&uac)
	require.Nil(t, err, "UpdateAPIClient() returned an error")
	require.NotNil(t, ac, "UpdateAPIClient() returned nil")
	assert.Equal(t, "cli-dsse2", ac.ID)
}

func TestDestroyAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/api_clients/cli-dsse2",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.DestroyAPIClient("cli-dsse2")
	require.Nil(t, err, "DestroyAPIClient() returned an error")
}

func TestResetSecretForAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients/cli-dsse2/reset_secret",
		ExpectBody:   "",
		GiveBody:     readJSON("api_client"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	ac, err := client.ResetSecretForAPIClient("cli-dsse2")
	require.Nil(t, err, "ResetPasswordForAPIClient() returned an error")
	require.NotNil(t, ac, "ResetPasswordForAPIClient() returned nil")
}

func TestLockAPIClient(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/api_clients/cli-dsse2/lock_resource",
		ExpectBody:   ``,
		GiveBody:     ``,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := gobrightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.LockResource(gobrightbox.APIClient{ID: "cli-dsse2"})
	require.Nil(t, err, "LockAPIClient() returned an error")
}

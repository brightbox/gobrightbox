package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestDatabaseServers(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/database_servers",
		ExpectBody:   "",
		GiveBody:     readJSON("database_servers"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	p, err := client.DatabaseServers()
	require.Nil(t, err, "DatabaseServers() returned an error")
	require.NotNil(t, p, "DatabaseServers() returned nil")
	require.Equal(t, 1, len(p), "wrong number of database servers returned")
	dbs := p[0]
	assert.Equal(t, "dbs-123ab", dbs.Id, "database server id incorrect")
}

func TestDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/database_servers/dbs-123ab",
		ExpectBody:   "",
		GiveBody:     readJSON("database_server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	dbs, err := client.DatabaseServer("dbs-123ab")
	require.Nil(t, err, "DatabaseServer() returned an error")
	require.NotNil(t, dbs, "DatabaseServer() returned nil")
	assert.Equal(t, "dbs-123ab", dbs.Id, "database server id incorrect")
	require.Equal(t, 2, len(dbs.AllowAccess), "not enough entries in the allow access list")
	stype := dbs.DatabaseServerType
	assert.Equal(t, "dbt-12345", stype.Id, "database server type Id incorrect")
}

func TestCreateDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/database_servers",
		ExpectBody:   "{}",
		GiveBody:     readJSON("database_server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	newDBS := brightbox.DatabaseServerOptions{}
	dbs, err := client.CreateDatabaseServer(&newDBS)
	require.Nil(t, err, "CreateDatabaseServer() returned an error")
	require.NotNil(t, dbs, "CreateDatabaseServer() returned nil")
	assert.Equal(t, "dbs-123ab", dbs.Id)
}

func TestCreateDatabaseServerWithAllowAccess(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/database_servers",
		ExpectBody:   `{"allow_access":["1.2.3.4","5.6.7.8"]}`,
		GiveBody:     readJSON("database_server_with_password"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	access := []string{"1.2.3.4", "5.6.7.8"}
	newDBS := brightbox.DatabaseServerOptions{AllowAccess: &access}
	dbs, err := client.CreateDatabaseServer(&newDBS)
	require.Nil(t, err, "CreateDatabaseServer() returned an error")
	require.NotNil(t, dbs, "CreateDatabaseServer() returned nil")
	assert.Equal(t, "dbs-123ab", dbs.Id)
}

func TestUpdateDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/database_servers/dbs-123ab",
		ExpectBody:   `{"name":"db server"}`,
		GiveBody:     readJSON("database_server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	name := "db server"
	udbs := brightbox.DatabaseServerOptions{Id: "dbs-123ab", Name: &name}
	dbs, err := client.UpdateDatabaseServer(&udbs)
	require.Nil(t, err, "UpdateDatabaseServer() returned an error")
	require.NotNil(t, dbs, "UpdateDatabaseServer() returned nil")
	assert.Equal(t, "dbs-123ab", dbs.Id)
}

func TestDestroyDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/database_servers/dbs-123ab",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.DestroyDatabaseServer("dbs-123ab")
	require.Nil(t, err, "DestroyDatabaseServer() returned an error")
}

func TestSnapshotDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/database_servers/dbs-123ab/snapshot",
		ExpectBody:   "",
		GiveBody:     readJSON("database_server"),
		GiveHeaders:   map[string]string{"Link": "<https://api.gb1.brightbox.com/1.0/database_snapshots/dbi-zlms8>; rel=snapshot"},
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	snap, err := client.SnapshotDatabaseServer("dbs-123ab")
	require.Nil(t, err, "SnapshotDatabaseServer() returned an error")
	require.NotNil(t, snap, "SnapshotDatabaseServer() returned nil")
	assert.Equal(t, "dbi-zlms8", snap.Id, "DatabaseSnapshot id incorrect")
}

func TestResetPasswordForDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/database_servers/dbs-123ab/reset_password",
		ExpectBody:   "",
		GiveBody:     readJSON("database_server_with_password"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	dbs, err := client.ResetPasswordForDatabaseServer("dbs-123ab")
	require.Nil(t, err, "ResetPasswordForDatabaseServer() returned an error")
	require.NotNil(t, dbs, "ResetPasswordForDatabaseServer() returned nil")
	assert.Equal(t, "admin", dbs.AdminUsername, "Database admin_username incorrect")
	assert.Equal(t, "k43;2kd432f", dbs.AdminPassword, "Database admin_password incorrect")
}

func TestLockDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/database_servers/dbs-aaaaa/lock_resource",
		ExpectBody:   ``,
		GiveBody:     ``,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	err = client.LockResource(brightbox.DatabaseServer{Id: "dbs-aaaaa"})
	require.Nil(t, err, "LockDatabaseServer() returned an error")
}

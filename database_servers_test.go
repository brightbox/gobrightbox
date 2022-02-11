package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
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
	assert.NilError(t, err, "NewClient returned an error")

	p, err := client.DatabaseServers()
	assert.NilError(t, err, "DatabaseServers() returned an error")
	assert.Assert(t, p != nil, "DatabaseServers() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of database servers returned")
	dbs := p[0]
	assert.Check(t, is.Equal("dbs-123ab", dbs.ID), "database server id incorrect")
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
	assert.NilError(t, err, "NewClient returned an error")

	dbs, err := client.DatabaseServer("dbs-123ab")
	assert.NilError(t, err, "DatabaseServer() returned an error")
	assert.Assert(t, dbs != nil, "DatabaseServer() returned nil")
	assert.Check(t, is.Equal("dbs-123ab", dbs.ID), "database server id incorrect")
	assert.Equal(t, 2, len(dbs.AllowAccess), "not enough entries in the allow access list")
	stype := dbs.DatabaseServerType
	assert.Check(t, is.Equal("dbt-12345", stype.ID), "database server type ID incorrect")
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
	assert.NilError(t, err, "NewClient returned an error")

	newDBS := brightbox.DatabaseServerOptions{}
	dbs, err := client.CreateDatabaseServer(&newDBS)
	assert.NilError(t, err, "CreateDatabaseServer() returned an error")
	assert.Assert(t, dbs != nil, "CreateDatabaseServer() returned nil")
	assert.Check(t, is.Equal("dbs-123ab", dbs.ID))
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
	assert.NilError(t, err, "NewClient returned an error")

	access := []string{"1.2.3.4", "5.6.7.8"}
	newDBS := brightbox.DatabaseServerOptions{AllowAccess: access}
	dbs, err := client.CreateDatabaseServer(&newDBS)
	assert.NilError(t, err, "CreateDatabaseServer() returned an error")
	assert.Assert(t, dbs != nil, "CreateDatabaseServer() returned nil")
	assert.Check(t, is.Equal("dbs-123ab", dbs.ID))
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
	assert.NilError(t, err, "NewClient returned an error")

	name := "db server"
	udbs := brightbox.DatabaseServerOptions{ID: "dbs-123ab", Name: &name}
	dbs, err := client.UpdateDatabaseServer(&udbs)
	assert.NilError(t, err, "UpdateDatabaseServer() returned an error")
	assert.Assert(t, dbs != nil, "UpdateDatabaseServer() returned nil")
	assert.Check(t, is.Equal("dbs-123ab", dbs.ID))
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
	assert.NilError(t, err, "NewClient returned an error")

	err = client.DestroyDatabaseServer("dbs-123ab")
	assert.NilError(t, err, "DestroyDatabaseServer() returned an error")
}

func TestSnapshotDatabaseServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/database_servers/dbs-123ab/snapshot",
		ExpectBody:   "",
		GiveBody:     readJSON("database_server"),
		GiveHeaders:  map[string]string{"Link": "<https://api.gb1.brightbox.com/1.0/database_snapshots/dbi-zlms8>; rel=snapshot"},
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	snap, err := client.SnapshotDatabaseServer("dbs-123ab")
	assert.NilError(t, err, "SnapshotDatabaseServer() returned an error")
	assert.Assert(t, snap != nil, "SnapshotDatabaseServer() returned nil")
	assert.Check(t, is.Equal("dbi-zlms8", snap.ID), "DatabaseSnapshot id incorrect")
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
	assert.NilError(t, err, "NewClient returned an error")

	dbs, err := client.ResetPasswordForDatabaseServer("dbs-123ab")
	assert.NilError(t, err, "ResetPasswordForDatabaseServer() returned an error")
	assert.Assert(t, dbs != nil, "ResetPasswordForDatabaseServer() returned nil")
	assert.Check(t, is.Equal("admin", dbs.AdminUsername), "Database admin_username incorrect")
	assert.Check(t, is.Equal("k43;2kd432f", dbs.AdminPassword), "Database admin_password incorrect")
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
	assert.NilError(t, err, "NewClient returned an error")

	err = client.LockResource(brightbox.DatabaseServer{ID: "dbs-aaaaa"})
	assert.NilError(t, err, "LockDatabaseServer() returned an error")
}

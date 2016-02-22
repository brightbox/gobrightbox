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


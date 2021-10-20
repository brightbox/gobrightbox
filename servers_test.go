package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServers(t *testing.T) {

	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/servers",
		ExpectBody:   ``,
		GiveBody:     readJSON("servers"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	servers, err := client.Servers()
	if err != nil {
		t.Fatal(err)
	}
	if len(servers) != 1 {
		t.Fatal("Wrong number of servers returned")
	}
	cs := servers
	s := cs[0]
	if s.ID != "srv-lv426" {
		t.Errorf("server ID incorrect")
	}
}

func TestServer(t *testing.T) {

	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/servers/srv-lv426",
		ExpectBody:   ``,
		GiveBody:     readJSON("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	s, err := client.Server("srv-lv426")
	if err != nil {
		t.Fatal(err)
	}
	if s.ID != "srv-lv426" {
		t.Error("server ID incorrect")
	}

	if !s.DiskEncrypted {
		t.Errorf("server DiskEncrypted was %v, should be true", s.DiskEncrypted)
	}
	if s.DeletedAt != nil {
		t.Errorf("server DeletedAt was %v, should be nil", s.DeletedAt)
	}
	if len(s.ServerGroups) != 1 || s.ServerGroups[0].ID != "grp-sda44" {
		t.Errorf("server groups is %v", s.ServerGroups)
	}
	if s.Image.ID != "img-3ikco" {
		t.Errorf("image is %q", s.Image.ID)
	}
	if s.Account.ID != "acc-43ks4" {
		t.Errorf("account is %q", s.Account.ID)
	}
}

func TestCreateServerWithImage(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/servers",
		ExpectBody:   map[string]string{"image": "img-12345"},
		GiveBody:     readJSON("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	opts := brightbox.ServerOptions{Image: "img-12345"}
	s, err := client.CreateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.ID != "srv-lv426" {
		t.Errorf("server ID is %s", s.ID)
	}

}

func TestCreateServerWithEmptyGroupList(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/servers",
		ExpectBody: map[string]interface{}{
			"image": "img-12345",
		},
		GiveBody: readJSON("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	groups := []string{}
	opts := brightbox.ServerOptions{
		Image:        "img-12345",
		ServerGroups: groups,
	}
	s, err := client.CreateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.ID != "srv-lv426" {
		t.Errorf("server ID is %s", s.ID)
	}

}

func TestCreateServerWithOptionalFields(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/servers",
		ExpectBody: map[string]interface{}{
			"image":              "img-12345",
			"name":               "myserver",
			"server_groups":      []string{"grp-aaaaa", "grp-bbbbb"},
			"compatibility_mode": true,
		},
		GiveBody: readJSON("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "myserver"
	groups := []string{"grp-aaaaa", "grp-bbbbb"}
	compat := true
	opts := brightbox.ServerOptions{
		Image:             "img-12345",
		Name:              &name,
		ServerGroups:      groups,
		CompatibilityMode: &compat,
	}
	s, err := client.CreateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.ID != "srv-lv426" {
		t.Errorf("server ID is %s", s.ID)
	}

}

func TestUpdateServerWithEmptyGroupsList(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/servers/srv-lv426",
		ExpectBody:   map[string]string{},
		GiveBody:     readJSON("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	groups := []string{}
	opts := brightbox.ServerOptions{ID: "srv-lv426", ServerGroups: groups}
	s, err := client.UpdateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.ID != "srv-lv426" {
		t.Errorf("server ID is %s", s.ID)
	}

}

func TestLockServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/servers/srv-lv426/lock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	s := new(brightbox.Server)
	s.ID = "srv-lv426"
	err = client.LockResource(s)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnLockServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/servers/srv-lv426/unlock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	s := new(brightbox.Server)
	s.ID = "srv-lv426"
	err = client.UnLockResource(s)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSnapshotServer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/servers/srv-aaaaa/snapshot",
		ExpectBody:   "",
		GiveBody:     readJSON("server"),
		GiveHeaders:  map[string]string{"Link": "<https://api.gb1.brightbox.com/1.0/images/img-bbbbb>; rel=snapshot"},
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	require.Nil(t, err, "NewClient returned an error")

	snap, err := client.SnapshotServer("srv-aaaaa")
	require.Nil(t, err, "SnapshotServer() returned an error")
	require.NotNil(t, snap, "SnapshotServer() returned nil")
	assert.Equal(t, "img-bbbbb", snap.ID, "Image id incorrect")
}

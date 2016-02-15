package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"testing"
)

func TestServers(t *testing.T) {

	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/servers",
		ExpectBody:   ``,
		GiveBody:     readJson("servers"),
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
	if s.Id != "srv-lv426" {
		t.Errorf("server Id incorrect")
	}
}

func TestServer(t *testing.T) {

	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/servers/srv-lv426",
		ExpectBody:   ``,
		GiveBody:     readJson("server"),
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
	if s.Id != "srv-lv426" {
		t.Error("server Id incorrect")
	}

	if s.DeletedAt != nil {
		t.Errorf("server DeletedAt was %v, should be nil", s.DeletedAt)
	}
	if len(s.ServerGroups) != 1 || s.ServerGroups[0].Id != "grp-sda44" {
		t.Errorf("server groups is %v", s.ServerGroups)
	}
	if s.Image.Id != "img-3ikco" {
		t.Errorf("image is %q", s.Image.Id)
	}
	if s.Account.Id != "acc-43ks4" {
		t.Errorf("account is %q", s.Account.Id)
	}
}

func TestCreateServerWithImage(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/servers",
		ExpectBody:   map[string]string{"image": "img-12345"},
		GiveBody:     readJson("server"),
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
	if s.Id != "srv-lv426" {
		t.Errorf("server Id is %s", s.Id)
	}

}

func TestCreateServerWithOptionalFields(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/servers",
		ExpectBody: map[string]interface{}{
			"image":              "img-12345",
			"name":               "myserver",
			"server_groups":      []string{"grp-aaaaa", "grp-bbbbb"},
			"compatibility_mode": true,
		},
		GiveBody: readJson("server"),
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
		ServerGroups:      &groups,
		CompatibilityMode: &compat,
	}
	s, err := client.CreateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.Id != "srv-lv426" {
		t.Errorf("server Id is %s", s.Id)
	}

}

func TestUpdateServerWithEmptyGroupsList(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/servers/srv-lv426",
		ExpectBody:   map[string]string{"server_groups": "[]"},
		GiveBody:     readJson("server"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	groups := []string{}
	opts := brightbox.ServerOptions{Id: "srv-lv426", ServerGroups: &groups}
	s, err := client.UpdateServer(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Server")
	}
	if s.Id != "srv-lv426" {
		t.Errorf("server Id is %s", s.Id)
	}

}


func TestLockServer(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/servers/srv-lv426/lock_resource",
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
	s.Id = "srv-lv426"
	err = client.LockResource(s)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnLockServer(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/servers/srv-lv426/unlock_resource",
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
	s.Id = "srv-lv426"
	err = client.UnLockResource(s)
	if err != nil {
		t.Fatal(err)
	}
}

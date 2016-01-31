package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"testing"
)

func TestCloudIPs(t *testing.T) {

	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/cloud_ips",
		ExpectBody:   ``,
		GiveBody:     readJson("cloud_ips"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	cips, err := client.CloudIPs()
	if err != nil {
		t.Fatal(err)
	}
	if len(cips) != 1 {
		t.Fatal("Wrong number of cloud ips returned")
	}
	cs := cips
	s := cs[0]
	if s.Id != "cip-k4a25" {
		t.Errorf("cloud Id incorrect")
	}
	if len(s.PortTranslators) != 1 {
		t.Fatal("cloud ip port translators incorrect")
	}
	pt := s.PortTranslators[0]
	if pt.Protocol != "http" {
		t.Errorf("cloud ip port translator protocol incorrect")
	}
	if pt.Incoming != 443 {
		t.Errorf("cloud ip port translator incoming port incorrect")
	}
	if pt.Outgoing != 2443 {
		t.Errorf("cloud ip port translator outgoing port incorrect")
	}
}

func TestCreateCloudIP(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/cloud_ips",
		ExpectBody:   `{"name":"product website ip"}`,
		GiveBody:     readJson("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "product website ip"
	opts := brightbox.CloudIPOptions{Name: &name}
	s, err := client.CreateCloudIP(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Cloud IP")
	}
	if s.Id != "cip-k4a25" {
		t.Errorf("cloud ip id is %s", s.Id)
	}

}

func TestCreateCloudIPWithPortTranslator(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/cloud_ips",
		ExpectBody:   `{"port_translators":[{"incoming":443,"outgoing":2443,"protocol":"tcp"}]}`,
		GiveBody:     readJson("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	pt := brightbox.PortTranslator {
		Incoming: 443,
		Outgoing: 2443,
		Protocol: "tcp",
	}
	opts := brightbox.CloudIPOptions{
		PortTranslators: &[]brightbox.PortTranslator{
			pt,
		},
	}
	s, err := client.CreateCloudIP(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Errorf("Didn't return a Cloud IP")
	}
	if s.Id != "cip-k4a25" {
		t.Errorf("cloud ip id is %s", s.Id)
	}

}


func TestUpdateCloudIP(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/cloud_ips/cip-k4a25",
		ExpectBody:   `{"name":"product website ip"}`,
		GiveBody:     readJson("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "product website ip"
	opts := brightbox.CloudIPOptions{Id: "cip-k4a25", Name: &name}
	cip, err := client.UpdateCloudIP(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if cip == nil {
		t.Errorf("Didn't return a Cloud IP")
	}
}

func TestUpdateCloudIPPostTranslator(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/cloud_ips/cip-k4a25",
		ExpectBody:   `{"port_translators":[{"incoming":443,"outgoing":2443,"protocol":"tcp"}]}`,
		GiveBody:     readJson("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	pt := brightbox.PortTranslator {
		Incoming: 443,
		Outgoing: 2443,
		Protocol: "tcp",
	}
	opts := brightbox.CloudIPOptions{
		Id: "cip-k4a25",
		PortTranslators: &[]brightbox.PortTranslator{
			pt,
		},
	}
	cip, err := client.UpdateCloudIP(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if cip == nil {
		t.Errorf("Didn't return a Cloud IP")
	}
}

func TestLockCloudIP(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/cloud_ips/cip-k4a25/lock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	cip := new(brightbox.CloudIP)
	cip.Id = "cip-k4a25"
	err = client.LockResource(cip)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnLockCloudIP(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/cloud_ips/cip-k4a25/unlock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	cip := new(brightbox.CloudIP)
	cip.Id = "cip-k4a25"
	err = client.UnLockResource(cip)
	if err != nil {
		t.Fatal(err)
	}
}

package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCloudIPs(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/cloud_ips",
		ExpectBody:   "",
		GiveBody:     readJSON("cloud_ips"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	cips, err := client.CloudIPs()
	assert.NilError(t, err, "CloudIPs() returned an error")
	assert.Equal(t, 1, len(cips), "Wrong number of cloud ips returned")
	cip := cips[0]
	assert.Check(t, is.Equal("cip-k4a25", cip.ID), "id doesn't match")
	assert.Equal(t, 1, len(cip.PortTranslators), "port translators list")
	pt := cip.PortTranslators[0]
	assert.Check(t, is.Equal("http", pt.Protocol), "port translator protocol")
	assert.Check(t, is.Equal(443, pt.Incoming), "port translator incoming port")
	assert.Check(t, is.Equal(2443, pt.Outgoing), "port translator outgoing port")
}

func TestCreateCloudIP(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/cloud_ips",
		ExpectBody:   map[string]string{"name": "product website ip"},
		GiveBody:     readJSON("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := "product website ip"
	opts := brightbox.CloudIPOptions{Name: &name}
	cip, err := client.CreateCloudIP(&opts)
	assert.NilError(t, err, "CloudIP() returned an error")
	assert.Assert(t, cip != nil, "didn't return a cloud ip")
	assert.Check(t, is.Equal("cip-k4a25", cip.ID), "cloud ip id")
}

func TestCreateCloudIPWithPortTranslator(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/cloud_ips",
		ExpectBody:   `{"port_translators":[{"incoming":443,"outgoing":2443,"protocol":"tcp"}]}`,
		GiveBody:     readJSON("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	pt := brightbox.PortTranslator{
		Incoming: 443,
		Outgoing: 2443,
		Protocol: "tcp",
	}
	opts := brightbox.CloudIPOptions{
		PortTranslators: []brightbox.PortTranslator{
			pt,
		},
	}
	cip, err := client.CreateCloudIP(&opts)
	assert.NilError(t, err, "CreateCloudIP returned an error")
	assert.Assert(t, cip != nil, "Didn't return a Cloud IP")
	assert.Check(t, is.Equal("cip-k4a25", cip.ID), "Cloud IP id didn't match")

}

func TestUpdateCloudIP(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/cloud_ips/cip-k4a25",
		ExpectBody:   map[string]string{"name": "product website ip"},
		GiveBody:     readJSON("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := "product website ip"
	opts := brightbox.CloudIPOptions{ID: "cip-k4a25", Name: &name}
	cip, err := client.UpdateCloudIP(&opts)
	assert.NilError(t, err, "NewClient returned an error")
	assert.Assert(t, cip != nil, "Didn't return a Cloud IP")
}

func TestUpdateCloudIPPortTranslator(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/cloud_ips/cip-k4a25",
		ExpectBody:   `{"port_translators":[{"incoming":443,"outgoing":2443,"protocol":"tcp"}]}`,
		GiveBody:     readJSON("cloud_ip"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	pt := brightbox.PortTranslator{
		Incoming: 443,
		Outgoing: 2443,
		Protocol: "tcp",
	}
	opts := brightbox.CloudIPOptions{
		ID: "cip-k4a25",
		PortTranslators: []brightbox.PortTranslator{
			pt,
		},
	}
	cip, err := client.UpdateCloudIP(&opts)
	assert.NilError(t, err, "UpdateCloudIP returned an error")
	assert.Assert(t, cip != nil, "UpdateCloudIP didn't return a cloud ip")
}

func TestLockCloudIP(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/cloud_ips/cip-k4a25/lock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	cip := new(brightbox.CloudIP)
	cip.ID = "cip-k4a25"
	err = client.LockResource(cip)
	assert.Assert(t, err != nil, "LockResource should return an error")
}

func TestUnLockCloudIP(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/cloud_ips/cip-k4a25/unlock_resource",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	cip := new(brightbox.CloudIP)
	cip.ID = "cip-k4a25"
	err = client.UnLockResource(cip)
	assert.Assert(t, err != nil, "UnLockResource should return an error")
}

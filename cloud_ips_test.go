package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
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
	require.Nil(t, err, "NewClient returned an error")

	cips, err := client.CloudIPs()
	require.Nil(t, err, "CloudIPs() returned an error")
	require.Equal(t, 1, len(cips), "Wrong number of cloud ips returned")
	cip := cips[0]
	assert.Equal(t, "cip-k4a25", cip.Id, "id doesn't match")
	require.Equal(t, 1, len(cip.PortTranslators), "port translators list")
	pt := cip.PortTranslators[0]
	assert.Equal(t, "http", pt.Protocol, "port translator protocol")
	assert.Equal(t, 443, pt.Incoming, "port translator incoming port")
	assert.Equal(t, 2443, pt.Outgoing, "port translator outgoing port")
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
	require.Nil(t, err, "NewClient returned an error")

	name := "product website ip"
	opts := brightbox.CloudIPOptions{Name: &name}
	cip, err := client.CreateCloudIP(&opts)
	require.Nil(t, err, "CloudIP() returned an error")
	require.NotNil(t, cip, "didn't return a cloud ip")
	assert.Equal(t, "cip-k4a25", cip.Id, "cloud ip id")
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
	require.Nil(t, err, "NewClient returned an error")

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
	require.Nil(t, err, "CreateCloudIP returned an error")
	require.NotNil(t, cip, "Didn't return a Cloud IP")
	assert.Equal(t, "cip-k4a25", cip.Id, "Cloud IP id didn't match")

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
	require.Nil(t, err, "NewClient returned an error")

	name := "product website ip"
	opts := brightbox.CloudIPOptions{Id: "cip-k4a25", Name: &name}
	cip, err := client.UpdateCloudIP(&opts)
	require.Nil(t, err, "NewClient returned an error")
	require.NotNil(t, cip, "Didn't return a Cloud IP")
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
	require.Nil(t, err, "NewClient returned an error")

	pt := brightbox.PortTranslator{
		Incoming: 443,
		Outgoing: 2443,
		Protocol: "tcp",
	}
	opts := brightbox.CloudIPOptions{
		Id: "cip-k4a25",
		PortTranslators: []brightbox.PortTranslator{
			pt,
		},
	}
	cip, err := client.UpdateCloudIP(&opts)
	require.Nil(t, err, "UpdateCloudIP returned an error")
	require.NotNil(t, cip, "UpdateCloudIP didn't return a cloud ip")
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
	require.Nil(t, err, "NewClient returned an error")

	cip := new(brightbox.CloudIP)
	cip.Id = "cip-k4a25"
	err = client.LockResource(cip)
	assert.Nil(t, err, "LockResource returned an error")
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
	require.Nil(t, err, "NewClient returned an error")

	cip := new(brightbox.CloudIP)
	cip.Id = "cip-k4a25"
	err = client.UnLockResource(cip)
	assert.Nil(t, err, "UnLockResource returned an error")
}

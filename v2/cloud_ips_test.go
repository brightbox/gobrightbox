package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestCloudIPs(t *testing.T) {
	instance := testAll[CloudIP](
		t,
		"CloudIP",
		"cloud_ips",
		"config map",
	)
	assert.Equal(t, instance.ID, "cip-k4a25")
}

func TestCloudIP(t *testing.T) {
	instance := testInstance[CloudIP](
		t,
		"CloudIP",
		"cloud_ips",
		"cloud_ip",
		"cip-k4a25",
	)
	assert.Equal(t, instance.Name, "product website ip")
}

func TestCreateCloudIP(t *testing.T) {
	newAC := CloudIPOptions{}
	_ = testCreate[CloudIP](
		t,
		"CloudIP",
		"cloud_ips",
		"cloud_ip",
		"cip-k4a25",
		&newAC,
		"{}",
	)
}

func TestUpdateCloudIP(t *testing.T) {
	name := "dev client"
	uac := CloudIPOptions{ID: "cip-k4a25", Name: &name}
	_ = testUpdate[CloudIP](
		t,
		"CloudIP",
		"cloud_ips",
		"cloud_ip",
		"cip-k4a25",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyCloudIP(t *testing.T) {
	testDestroy[CloudIP](
		t,
		"CloudIP",
		"cloud_ips",
		"cip-k4a25",
	)
}

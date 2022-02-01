package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestServerGroups(t *testing.T) {
	instance := testAll[ServerGroup](
		t,
		"ServerGroup",
		"server_groups",
		"server group",
	)
	assert.Equal(t, instance.ID, "grp-sda44")
}

func TestServerGroup(t *testing.T) {
	instance := testInstance[ServerGroup](
		t,
		"ServerGroup",
		"server_groups",
		"server_group",
		"grp-sda44",
	)
	assert.Equal(t, instance.Name, "default")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.FirewallPolicy.ID, "fwp-j3654")
}

func TestCreateServerGroup(t *testing.T) {
	newAC := ServerGroupOptions{}
	_ = testCreate[ServerGroup](
		t,
		"ServerGroup",
		"server_groups",
		"server_group",
		"grp-sda44",
		&newAC,
		"{}",
	)
}

func TestUpdateServerGroup(t *testing.T) {
	name := "default"
	uac := ServerGroupOptions{ID: "grp-sda44", Name: &name}
	_ = testUpdate[ServerGroup](
		t,
		"ServerGroup",
		"server_groups",
		"server_group",
		"grp-sda44",
		&uac,
		`{"name":"default"}`,
	)
}

func TestDestroyServerGroup(t *testing.T) {
	testDestroy[ServerGroup](
		t,
		"ServerGroup",
		"server_groups",
		"grp-sda44",
	)
}

package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestLoadBalancers(t *testing.T) {
	instance := testAll[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		"load_balancer",
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestLoadBalancer(t *testing.T) {
	instance := testInstance[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		"load_balancer",
		"lba-1235f",
	)
	assert.Equal(t, instance.Status, "creating")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
	assert.Equal(t, instance.Nodes[0].ID, "srv-lv426")
	assert.Assert(t, is.Len(instance.CloudIPs, 0))
}

func TestCreateLoadBalancer(t *testing.T) {
	newAC := LoadBalancerOptions{}
	_ = testCreate[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		"load_balancer",
		"lba-1235f",
		&newAC,
		"{}",
	)
}

func TestUpdateLoadBalancer(t *testing.T) {
	name := "dev client"
	uac := LoadBalancerOptions{ID: "lba-1235f", Name: &name}
	_ = testUpdate[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		"load_balancer",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyLoadBalancer(t *testing.T) {
	testDestroy[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		"lba-1235f",
	)
}

func TestLockLoadBalancer(t *testing.T) {
	testLock[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		&LoadBalancer{ID: "lba-1235f"},
		"lock_resource",
		LockResource,
	)
}

func TestUnlockLoadBalancer(t *testing.T) {
	testLock[LoadBalancer](
		t,
		"LoadBalancer",
		"load_balancers",
		&LoadBalancer{ID: "lba-1235f"},
		"unlock_resource",
		UnlockResource,
	)
}

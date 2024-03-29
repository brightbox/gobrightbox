// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestLoadBalancers(t *testing.T) {
	instance := testAll(
		t,
		(*Client).LoadBalancers,
		"LoadBalancer",
		"load_balancers",
		"LoadBalancers",
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestLoadBalancer(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).LoadBalancer,
		"LoadBalancer",
		path.Join("load_balancers", "lba-1235f"),
		"load_balancer",
		"lba-1235f",
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestCreateLoadBalancer(t *testing.T) {
	newResource := LoadBalancerOptions{}
	instance := testModify(
		t,
		(*Client).CreateLoadBalancer,
		newResource,
		"load_balancer",
		"POST",
		path.Join("load_balancers"),
		"{}",
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestUpdateLoadBalancer(t *testing.T) {
	updatedResource := LoadBalancerOptions{ID: "lba-1235f"}
	instance := testModify(
		t,
		(*Client).UpdateLoadBalancer,
		updatedResource,
		"load_balancer",
		"PUT",
		path.Join("load_balancers", updatedResource.ID),
		"{}",
	)
	assert.Equal(t, instance.ID, updatedResource.ID)
}

func TestDestroyLoadBalancer(t *testing.T) {
	deletedResource := testModify(
		t,
		(*Client).DestroyLoadBalancer,
		"lba-1235f",
		"load_balancer",
		"DELETE",
		path.Join("load_balancers", "lba-1235f"),
		"",
	)
	assert.Equal(t, deletedResource.ID, "lba-1235f")
}

func TestLockLoadBalancer(t *testing.T) {
	lockedResource := testModify(
		t,
		(*Client).LockLoadBalancer,
		"lba-1235f",
		"load_balancer",
		"PUT",
		path.Join("load_balancers", "lba-1235f", "lock_resource"),
		"",
	)
	assert.Equal(t, lockedResource.ID, "lba-1235f")
}

func TestUnlockLoadBalancer(t *testing.T) {
	unlockedResource := testModify(
		t,
		(*Client).UnlockLoadBalancer,
		"lba-1235f",
		"load_balancer",
		"PUT",
		path.Join("load_balancers", "lba-1235f", "unlock_resource"),
		"",
	)
	assert.Equal(t, unlockedResource.ID, "lba-1235f")
}

func TestLoadBalancerCreatedAtUnix(t *testing.T) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	target := LoadBalancer{CreatedAt: &tm}
	assert.Equal(t, target.CreatedAtUnix(), tm.Unix())
}

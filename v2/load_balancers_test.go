package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestApplyLoadBalancer(t *testing.T) {
	instance := testLink[LoadBalancer, []LoadBalancerNode](
		t,
		(*Client).AddNodesToLoadBalancer,
		"lba-1235f",
		[]LoadBalancerNode{{"srv-lv426"}},
		"load_balancer",
		"POST",
		path.Join("load_balancers", "lba-1235f", "add_nodes"),
		`[{"node":"srv-lv426"}]`,
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestRemoveLoadBalancer(t *testing.T) {
	instance := testLink[LoadBalancer, []LoadBalancerNode](
		t,
		(*Client).RemoveNodesFromLoadBalancer,
		"lba-1235f",
		[]LoadBalancerNode{{"srv-lv426"}},
		"load_balancer",
		"POST",
		path.Join("load_balancers", "lba-1235f", "remove_nodes"),
		`[{"node":"srv-lv426"}]`,
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

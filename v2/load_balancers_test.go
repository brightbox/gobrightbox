package brightbox

import (
	"path"
	"testing"

	"github.com/brightbox/gobrightbox/v2/status/listenerprotocol"
	"github.com/brightbox/gobrightbox/v2/status/proxyprotocol"

	"gotest.tools/v3/assert"
)

func TestAddNodesToLoadBalancer(t *testing.T) {
	instance := testLink(
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

func TestRemoveNodesFromLoadBalancer(t *testing.T) {
	instance := testLink(
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

func TestAddListenersToLoadBalancer(t *testing.T) {
	instance := testLink(
		t,
		(*Client).AddListenersToLoadBalancer,
		"lba-1235f",
		[]LoadBalancerListener{
			{
				Protocol: listenerprotocol.Http,
				In:       80,
				Out:      80,
				Timeout:  50000,
			},
		},
		"load_balancer",
		"POST",
		path.Join("load_balancers", "lba-1235f", "add_listeners"),
		`[{"protocol":"http","in":80,"out":80,"timeout":50000}]`,
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

func TestRemoveListenersFromLoadBalancer(t *testing.T) {
	instance := testLink(
		t,
		(*Client).RemoveListenersFromLoadBalancer,
		"lba-1235f",
		[]LoadBalancerListener{
			{
				Protocol:      listenerprotocol.Http,
				In:            80,
				Out:           80,
				Timeout:       50000,
				ProxyProtocol: proxyprotocol.V2SslCn,
			},
		},
		"load_balancer",
		"POST",
		path.Join("load_balancers", "lba-1235f", "remove_listeners"),
		`[{"protocol":"http","in":80,"out":80,"timeout":50000,"proxy_protocol":"v2-ssl-cn"}]`,
	)
	assert.Equal(t, instance.ID, "lba-1235f")
}

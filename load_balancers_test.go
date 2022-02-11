package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLoadBalancers(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/load_balancers",
		ExpectBody:   "",
		GiveBody:     readJSON("load_balancers"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	p, err := client.LoadBalancers()
	assert.NilError(t, err, "LoadBalancers() returned an error")
	assert.Assert(t, p != nil, "LoadBalancers() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of load balancers returned")
	lb := p[0]
	assert.Check(t, is.Equal("lba-1235f", lb.ID), "load balancer id incorrect")
	assert.Equal(t, 1, len(lb.Nodes), "not enough nodes returned")
	node := lb.Nodes[0]
	assert.Check(t, is.Equal("srv-lv426", node.ID), "node ID incorrect")
}

func TestLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/load_balancers/lba-1235f",
		ExpectBody:   "",
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	lb, err := client.LoadBalancer("lba-1235f")
	assert.NilError(t, err, "LoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "LoadBalancer() returned nil")
	assert.Check(t, is.Equal("lba-1235f", lb.ID), "load balancer id incorrect")
	assert.Equal(t, 1, len(lb.Nodes), "not enough nodes returned")

	node := lb.Nodes[0]
	assert.Check(t, is.Equal("srv-lv426", node.ID), "node ID incorrect")

	assert.Equal(t, 1, len(lb.Listeners), "not enough listeners")
	lnr := lb.Listeners[0]
	assert.Check(t, is.Equal(80, lnr.In), "listener in port incorrect")
	assert.Check(t, is.Equal(80, lnr.Out), "listener out port incorrect")
	assert.Check(t, is.Equal(50000, lnr.Timeout), "listener timeout incorrect")
	assert.Check(t, is.Equal("http", lnr.Protocol), "listener protocol incorrect")
	assert.Check(t, is.Len(lnr.ProxyProtocol, 0), "proxy protocol should be empty")

	assert.Check(t, is.Equal("http", lb.Healthcheck.Type), "healthcheck type incorrect")
	assert.Check(t, is.Equal("/", lb.Healthcheck.Request), "healthcheck request incorrect")
	assert.Check(t, is.Equal(80, lb.Healthcheck.Port), "healthchech port incorrect")

	assert.Equal(t, lb.HTTPSRedirect, false, "https redirect should be off")
	assert.Assert(t, lb.Certificate != nil, "certificate is nil")
	assert.Check(t, is.Equal("/CN=www.example.com", lb.Certificate.Subject), "certificate subject is incorrect")
}

func TestUpdateLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa",
		ExpectBody:   map[string]string{"name": "my lb"},
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := "my lb"
	newLB := brightbox.LoadBalancerOptions{ID: "lba-aaaaa", Name: &name}
	lb, err := client.UpdateLoadBalancer(&newLB)
	assert.NilError(t, err, "UpdateLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "UpdateLoadBalancer() returned nil")
}

func TestCreateLoadBalancerWithHealthCheck(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers",
		ExpectBody:   `{"healthcheck":{"type":"http","port":80,"request":"/health"}}`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	hc := brightbox.LoadBalancerHealthcheck{Type: "http", Port: 80, Request: "/health"}
	newLB := brightbox.LoadBalancerOptions{Healthcheck: &hc}
	lb, err := client.CreateLoadBalancer(&newLB)
	assert.NilError(t, err, "CreateLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "CreateLoadBalancer() returned nil")
}

func TestCreateLoadBalancerWithListeners(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers",
		ExpectBody:   `{"listeners":[{"protocol":"http","in":80,"out":8080}]}`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	l := brightbox.LoadBalancerListener{Protocol: "http", In: 80, Out: 8080}
	ls := []brightbox.LoadBalancerListener{l}
	newLB := brightbox.LoadBalancerOptions{Listeners: ls}
	lb, err := client.CreateLoadBalancer(&newLB)
	assert.NilError(t, err, "CreateLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "CreateLoadBalancer() returned nil")
}

func TestCreateLoadBalancerWithNodes(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers",
		ExpectBody:   `{"nodes":[{"node":"srv-aaaaa"}]}`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	n := brightbox.LoadBalancerNode{Node: "srv-aaaaa"}
	ns := []brightbox.LoadBalancerNode{n}
	newLB := brightbox.LoadBalancerOptions{Nodes: ns}
	lb, err := client.CreateLoadBalancer(&newLB)
	assert.NilError(t, err, "CreateLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "CreateLoadBalancer() returned nil")
}

func TestCreateLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers",
		ExpectBody:   "{}",
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	newLB := new(brightbox.LoadBalancerOptions)
	lb, err := client.CreateLoadBalancer(newLB)
	assert.NilError(t, err, "CreateLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "CreateLoadBalancer() returned nil")
}

func TestDestroyLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	err = client.DestroyLoadBalancer("lba-aaaaa")
	assert.NilError(t, err, "DestroyLoadBalancer() returned an error")
}

func TestAddNodesToLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/add_nodes",
		ExpectBody:   `[{"node":"srv-aaaaa"}]`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	nodes := []brightbox.LoadBalancerNode{brightbox.LoadBalancerNode{Node: "srv-aaaaa"}}
	lb, err := client.AddNodesToLoadBalancer("lba-aaaaa", nodes)
	assert.NilError(t, err, "AddNodesToLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "AddNodesToLoadBalancer() returned nil")
}

func TestRemoveNodesFromLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/remove_nodes",
		ExpectBody:   `[{"node":"srv-bbbbb"}]`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	nodes := []brightbox.LoadBalancerNode{brightbox.LoadBalancerNode{Node: "srv-bbbbb"}}
	lb, err := client.RemoveNodesFromLoadBalancer("lba-aaaaa", nodes)
	assert.NilError(t, err, "RemoveNodesFromLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "RemoveNodesFromLoadBalancer() returned nil")
}

func TestAddListenersToLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/add_listeners",
		ExpectBody:   `[{"protocol":"tcp","in":80}]`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	listeners := []brightbox.LoadBalancerListener{brightbox.LoadBalancerListener{Protocol: "tcp", In: 80}}
	lb, err := client.AddListenersToLoadBalancer("lba-aaaaa", listeners)
	assert.NilError(t, err, "AddListenersToLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "AddListenersToLoadBalancer() returned nil")
}

func TestRemoveListenersFromLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/remove_listeners",
		ExpectBody:   `[{"protocol":"tcp","out":8080}]`,
		GiveBody:     readJSON("load_balancer"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	listeners := []brightbox.LoadBalancerListener{brightbox.LoadBalancerListener{Protocol: "tcp", Out: 8080}}
	lb, err := client.RemoveListenersFromLoadBalancer("lba-aaaaa", listeners)
	assert.NilError(t, err, "RemoveListenersFromLoadBalancer() returned an error")
	assert.Assert(t, lb != nil, "RemoveListenersFromLoadBalancer() returned nil")
}

func TestLockLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/lock_resource",
		ExpectBody:   ``,
		GiveBody:     ``,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	err = client.LockResource(brightbox.LoadBalancer{ID: "lba-aaaaa"})
	assert.NilError(t, err, "LockLoadBalancer() returned an error")
}

func TestUnLockLoadBalancer(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/load_balancers/lba-aaaaa/unlock_resource",
		ExpectBody:   ``,
		GiveBody:     ``,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	err = client.UnLockResource(brightbox.LoadBalancer{ID: "lba-aaaaa"})
	assert.NilError(t, err, "UnLockLoadBalancer() returned an error")
}

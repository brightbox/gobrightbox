package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
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
	require.Nil(t, err, "NewClient returned an error")

	p, err := client.LoadBalancers()
	require.Nil(t, err, "LoadBalancers() returned an error")
	require.NotNil(t, p, "LoadBalancers() returned nil")
	require.Equal(t, 1, len(p), "wrong number of load balancers returned")
	lb := p[0]
	assert.Equal(t, "lba-1235f", lb.Id, "load balancer id incorrect")
	require.Equal(t, 1, len(lb.Nodes), "not enough nodes returned")
	node := lb.Nodes[0]
	assert.Equal(t, "srv-lv426", node.Id, "node Id incorrect")
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
	require.Nil(t, err, "NewClient returned an error")

	lb, err := client.LoadBalancer("lba-1235f")
	require.Nil(t, err, "LoadBalancer() returned an error")
	require.NotNil(t, lb, "LoadBalancer() returned nil")
	assert.Equal(t, "lba-1235f", lb.Id, "load balancer id incorrect")
	require.Equal(t, 1, len(lb.Nodes), "not enough nodes returned")

	node := lb.Nodes[0]
	assert.Equal(t, "srv-lv426", node.Id, "node Id incorrect")

	require.Equal(t, 1, len(lb.Listeners), "not enough listeners")
	lnr := lb.Listeners[0]
	assert.Equal(t, 80, lnr.In, "listener in port incorrect")
	assert.Equal(t, 80, lnr.Out, "listener out port incorrect")
	assert.Equal(t, 50000, lnr.Timeout, "listener timeout incorrect")
	assert.Equal(t, "http", lnr.Protocol, "listener protocol incorrect")

	assert.Equal(t, "http", lb.Healthcheck.Type, "healthcheck type incorrect")
	assert.Equal(t, "/", lb.Healthcheck.Request, "healthcheck request incorrect")
	assert.Equal(t, 80, lb.Healthcheck.Port, "healthchech port incorrect")

	require.NotNil(t, lb.Certificate, "certificate is nil")
	assert.Equal(t, "/CN=www.example.com", lb.Certificate.Subject, "certificate subject is incorrect")
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
	require.Nil(t, err, "NewClient returned an error")

	name := "my lb"
	newLB := brightbox.LoadBalancerOptions{Id: "lba-aaaaa", Name: &name}
	lb, err := client.UpdateLoadBalancer(&newLB)
	require.Nil(t, err, "UpdateLoadBalancer() returned an error")
	require.NotNil(t, lb, "UpdateLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	hc := brightbox.LoadBalancerHealthcheck{Type: "http", Port: 80, Request: "/health"}
	newLB := brightbox.LoadBalancerOptions{Healthcheck: &hc}
	lb, err := client.CreateLoadBalancer(&newLB)
	require.Nil(t, err, "CreateLoadBalancer() returned an error")
	require.NotNil(t, lb, "CreateLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	l := brightbox.LoadBalancerListener{Protocol: "http", In: 80, Out: 8080}
	ls := []brightbox.LoadBalancerListener{l}
	newLB := brightbox.LoadBalancerOptions{Listeners: &ls}
	lb, err := client.CreateLoadBalancer(&newLB)
	require.Nil(t, err, "CreateLoadBalancer() returned an error")
	require.NotNil(t, lb, "CreateLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	n := brightbox.LoadBalancerNode{Node: "srv-aaaaa"}
	ns := []brightbox.LoadBalancerNode{n}
	newLB := brightbox.LoadBalancerOptions{Nodes: &ns}
	lb, err := client.CreateLoadBalancer(&newLB)
	require.Nil(t, err, "CreateLoadBalancer() returned an error")
	require.NotNil(t, lb, "CreateLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	newLB := new(brightbox.LoadBalancerOptions)
	lb, err := client.CreateLoadBalancer(newLB)
	require.Nil(t, err, "CreateLoadBalancer() returned an error")
	require.NotNil(t, lb, "CreateLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	err = client.DestroyLoadBalancer("lba-aaaaa")
	require.Nil(t, err, "DestroyLoadBalancer() returned an error")
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
	require.Nil(t, err, "NewClient returned an error")

	nodes := []brightbox.LoadBalancerNode{brightbox.LoadBalancerNode{Node: "srv-aaaaa"}}
	lb, err := client.AddNodesToLoadBalancer("lba-aaaaa", nodes)
	require.Nil(t, err, "AddNodesToLoadBalancer() returned an error")
	require.NotNil(t, lb, "AddNodesToLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	nodes := []brightbox.LoadBalancerNode{brightbox.LoadBalancerNode{Node: "srv-bbbbb"}}
	lb, err := client.RemoveNodesFromLoadBalancer("lba-aaaaa", nodes)
	require.Nil(t, err, "RemoveNodesFromLoadBalancer() returned an error")
	require.NotNil(t, lb, "RemoveNodesFromLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	listeners := []brightbox.LoadBalancerListener{brightbox.LoadBalancerListener{Protocol: "tcp", In: 80}}
	lb, err := client.AddListenersToLoadBalancer("lba-aaaaa", listeners)
	require.Nil(t, err, "AddListenersToLoadBalancer() returned an error")
	require.NotNil(t, lb, "AddListenersToLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	listeners := []brightbox.LoadBalancerListener{brightbox.LoadBalancerListener{Protocol: "tcp", Out: 8080}}
	lb, err := client.RemoveListenersFromLoadBalancer("lba-aaaaa", listeners)
	require.Nil(t, err, "RemoveListenersFromLoadBalancer() returned an error")
	require.NotNil(t, lb, "RemoveListenersFromLoadBalancer() returned nil")
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
	require.Nil(t, err, "NewClient returned an error")

	err = client.LockResource(brightbox.LoadBalancer{Id: "lba-aaaaa"})
	require.Nil(t, err, "LockLoadBalancer() returned an error")
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
	require.Nil(t, err, "NewClient returned an error")

	err = client.UnLockResource(brightbox.LoadBalancer{Id: "lba-aaaaa"})
	require.Nil(t, err, "UnLockLoadBalancer() returned an error")
}

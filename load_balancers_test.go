package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestLoadBalancers(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/load_balancers",
		ExpectBody:   "",
		GiveBody:     readJson("load_balancers"),
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
	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/load_balancers/lba-1235f",
		ExpectBody:   "",
		GiveBody:     readJson("load_balancer"),
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
}

package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"testing"
)

func TestFirewallPolicies(t *testing.T) {

	handler := ApiMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectUrl:    "/1.0/firewall_policies",
		ExpectBody:   ``,
		GiveBody:     readJson("firewall_policies"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	policies, err := client.FirewallPolicies()
	if err != nil {
		t.Fatal(err)
	}
	if len(policies) != 1 {
		t.Fatal("Wrong number of policies returned")
	}
	p := policies[0]
	if p.Id != "fwp-j3654" {
		t.Errorf("policy id incorrect")
	}
	if p.ServerGroup.Name != "default" {
		t.Fatal("policy server group incorrect")
	}
	if len(p.Rules) != 1 {
		t.Errorf("policy rules incorrect")
	} else {
		if p.Rules[0].Id != "fwr-k32ls" {
			t.Errorf("policy rule id incorrext")
		}
	}
}

func TestCreateFirewallPolicy(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/firewall_policies",
		ExpectBody:   `{"name":"web servers"}`,
		GiveBody:     readJson("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "web servers"
	opts := brightbox.FirewallPolicyOptions{Name: &name}
	p, err := client.CreateFirewallPolicy(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Errorf("Didn't return a firewall policy")
	}
	if p.Id != "fwp-j3654" {
		t.Errorf("firewall policy id is %s", p.Id)
	}

}

func TestCreateFirewallPolicyWithServerGroup(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/firewall_policies",
		ExpectBody:   `{"name":"web servers","server_group":"grp-abcde"}`,
		GiveBody:     readJson("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "web servers"
	group := "grp-abcde"
	opts := brightbox.FirewallPolicyOptions{Name: &name, ServerGroup: &group}
	_, err = client.CreateFirewallPolicy(&opts)
	if err != nil {
		t.Fatal(err)
	}
}

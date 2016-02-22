package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"testing"
)

func TestFirewallPolicies(t *testing.T) {

	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/firewall_policies",
		ExpectBody:   ``,
		GiveBody:     readJSON("firewall_policies"),
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
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/firewall_policies",
		ExpectBody:   `{"name":"web servers"}`,
		GiveBody:     readJSON("firewall_policy"),
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
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/firewall_policies",
		ExpectBody:   `{"name":"web servers","server_group":"grp-abcde"}`,
		GiveBody:     readJSON("firewall_policy"),
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

func TestUpdateFirewallPolicy(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/firewall_policies/fwp-j3654",
		ExpectBody:   `{"name":"mail servers"}`,
		GiveBody:     readJSON("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	name := "mail servers"
	opts := brightbox.FirewallPolicyOptions{Id: "fwp-j3654", Name: &name}
	_, err = client.UpdateFirewallPolicy(&opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDestroyFirewallPolicy(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/firewall_policies/fwp-j3654",
		ExpectBody:   ``,
		GiveBody:     readJSON("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = client.DestroyFirewallPolicy("fwp-j3654")
	if err != nil {
		t.Fatal(err)
	}
}

func TestApplyFirewallPolicy(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/firewall_policies/fwp-j3654/apply_to",
		ExpectBody:   `{"server_group":"grp-abcde"}`,
		GiveBody:     readJSON("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ApplyFirewallPolicy("fwp-j3654", "grp-abcde")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFirewallPolicy(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/firewall_policies/fwp-j3654/remove",
		ExpectBody:   `{"server_group":"grp-abcde"}`,
		GiveBody:     readJSON("firewall_policy"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.RemoveFirewallPolicy("fwp-j3654", "grp-abcde")
	if err != nil {
		t.Fatal(err)
	}
}

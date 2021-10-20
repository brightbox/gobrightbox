package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
)

func TestCreateFirewallRule(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/firewall_rules",
		ExpectBody: map[string]string{
			"firewall_policy": "fwp-j3654",
			"protocol":        "tcp",
			"source":          "grp-xxxxx",
			"destination":     ""},
		GiveBody: readJSON("firewall_rule"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	pol := "fwp-j3654"
	proto := "tcp"
	dst := ""
	src := "grp-xxxxx"
	opts := brightbox.FirewallRuleOptions{
		FirewallPolicy: pol,
		Protocol:       &proto,
		Source:         &src,
		Destination:    &dst,
	}
	p, err := client.CreateFirewallRule(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Errorf("Didn't return a firewall rule")
	}
	if p.ID != "fwr-k32ls" {
		t.Errorf("firewall rule id is %s", p.ID)
	}

}

func TestUpdateFirewallRule(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/firewall_rules/fwr-k32ls",
		ExpectBody: map[string]string{
			"protocol":    "tcp",
			"source":      "grp-xxxxx",
			"destination": ""},
		GiveBody: readJSON("firewall_rule"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	proto := "tcp"
	dst := ""
	src := "grp-xxxxx"
	opts := brightbox.FirewallRuleOptions{
		ID:          "fwr-k32ls",
		Protocol:    &proto,
		Source:      &src,
		Destination: &dst,
	}
	p, err := client.UpdateFirewallRule(&opts)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Errorf("Didn't return a firewall rule")
	}
	if p.ID != "fwr-k32ls" {
		t.Errorf("firewall rule id is %s", p.ID)
	}
}

func TestDestroyFirewallRule(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/firewall_rules/fwr-j3654",
		ExpectBody:   ``,
		GiveBody:     readJSON("firewall_rule"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = client.DestroyFirewallRule("fwr-j3654")
	if err != nil {
		t.Fatal(err)
	}
}

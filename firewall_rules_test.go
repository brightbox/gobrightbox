package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"testing"
)

func TestCreateFirewallRule(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/1.0/firewall_rules",
		ExpectBody:   `{"firewall_policy":"fwp-j3654","protocol":"tcp","source":"grp-xxxxx","destination":""}`,
		GiveBody:     readJson("firewall_rule"),
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
	if p.Id != "fwr-k32ls" {
		t.Errorf("firewall rule id is %s", p.Id)
	}

}

func TestUpdateFirewallRule(t *testing.T) {
	handler := ApiMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectUrl:    "/1.0/firewall_rules/fwr-k32ls",
		ExpectBody:   map[string]string{
			"protocol":"tcp",
			"source":"grp-xxxxx",
			"destination":""},
		GiveBody:     readJson("firewall_rule"),
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
		Id: "fwr-k32ls",
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
	if p.Id != "fwr-k32ls" {
		t.Errorf("firewall rule id is %s", p.Id)
	}

}

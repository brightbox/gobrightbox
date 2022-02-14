package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestFirewallRules(t *testing.T) {
	instance := testAll[FirewallRule](
		t,
		"FirewallRule",
		"firewall_rules",
		"firewall rule",
	)
	assert.Equal(t, instance.ID, "fwr-k32ls")
}

func TestFirewallRule(t *testing.T) {
	instance := testInstance[FirewallRule](
		t,
		"FirewallRule",
		"firewall_rules",
		"firewall_rule",
		"fwr-k32ls",
	)
	assert.Equal(t, instance.Source, "srv-lv426")
	assert.Equal(t, instance.FirewallPolicy.ID, "fwp-j3654")
}

func TestCreateFirewallRule(t *testing.T) {
	newAC := FirewallRuleOptions{}
	_ = testCreate[FirewallRule](
		t,
		"FirewallRule",
		"firewall_rules",
		"firewall_rule",
		"fwr-k32ls",
		&newAC,
		"{}",
	)
}

func TestUpdateFirewallRule(t *testing.T) {
	name := "fwp-j3654"
	uac := FirewallRuleOptions{ID: "fwr-k32ls", FirewallPolicy: &name}
	_ = testUpdate[FirewallRule](
		t,
		"FirewallRule",
		"firewall_rules",
		"firewall_rule",
		&uac,
		`{"firewall_policy":"fwp-j3654"}`,
	)
}

func TestDestroyFirewallRule(t *testing.T) {
	testDestroy[FirewallRule](
		t,
		"FirewallRule",
		"firewall_rules",
		"fwr-k32ls",
	)
}

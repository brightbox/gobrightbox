package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestFirewallPolicies(t *testing.T) {
	instance := testAll[FirewallPolicy](
		t,
		"FirewallPolicy",
		"firewall_policies",
		"firewall policy",
	)
	assert.Equal(t, instance.ID, "fwp-j3654")
}

func TestFirewallPolicy(t *testing.T) {
	instance := testInstance[FirewallPolicy](
		t,
		"FirewallPolicy",
		"firewall_policies",
		"firewall_policy",
		"fwp-j3654",
	)
	assert.Equal(t, instance.Name, "default")
}

func TestCreateFirewallPolicy(t *testing.T) {
	newAC := FirewallPolicyOptions{}
	_ = testCreate[FirewallPolicy](
		t,
		"FirewallPolicy",
		"firewall_policies",
		"firewall_policy",
		"fwp-j3654",
		&newAC,
		"{}",
	)
}

func TestUpdateFirewallPolicy(t *testing.T) {
	name := "default"
	uac := FirewallPolicyOptions{ID: "fwp-j3654", Name: &name}
	_ = testUpdate[FirewallPolicy](
		t,
		"FirewallPolicy",
		"firewall_policies",
		"firewall_policy",
		"fwp-j3654",
		&uac,
		`{"name":"default"}`,
	)
}

func TestDestroyFirewallPolicy(t *testing.T) {
	testDestroy[FirewallPolicy](
		t,
		"FirewallPolicy",
		"firewall_policies",
		"fwp-j3654",
	)
}

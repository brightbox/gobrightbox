package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestApplyFirewallPolicy(t *testing.T) {
	instance := testLink[FirewallPolicy, FirewallPolicyAttachment](
		t,
		(*Client).ApplyFirewallPolicy,
		"fwp-j3654",
		FirewallPolicyAttachment{"grp-12345"},
		"firewall_policy",
		"POST",
		path.Join("firewall_policies", "fwp-j3654", "apply_to"),
		`{"server_group":"grp-12345"}`,
	)
	assert.Equal(t, instance.ID, "fwp-j3654")
}

func TestRemoveFirewallPolicy(t *testing.T) {
	instance := testLink[FirewallPolicy, FirewallPolicyAttachment](
		t,
		(*Client).RemoveFirewallPolicy,
		"fwp-j3654",
		FirewallPolicyAttachment{"grp-12345"},
		"firewall_policy",
		"POST",
		path.Join("firewall_policies", "fwp-j3654", "remove"),
		`{"server_group":"grp-12345"}`,
	)
	assert.Equal(t, instance.ID, "fwp-j3654")
}

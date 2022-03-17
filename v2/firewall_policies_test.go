
package brightbox

import (
	"context"
	"path"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func testLink[O, I any](
	t *testing.T,
	modify func(*Client, context.Context, string, I) (*O, error),
	from string,
	to I,
	jsonPath string,
	verb string,
	expectedPath string,
	expectedBody string,
) *O {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: verb,
			ExpectURL:    "/1.0/" + expectedPath,
			ExpectBody:   expectedBody,
			GiveBody:     readJSON(jsonPath),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")
	instance, err := modify(client, context.Background(), from, to)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, instance != nil)
	return instance
}

func TestApplyFirewallPolicy(t *testing.T) {
	instance := testLink[FirewallPolicy, FirewallPolicyGroup](
		t,
		(*Client).ApplyFirewallPolicy,
		"fwp-j3654",
		FirewallPolicyGroup{"grp-12345"},
		"firewall_policy",
		"POST",
		path.Join("firewall_policies", "fwp-j3654", "apply_to"),
		`{"server_group":"grp-12345"}`,
	)
	assert.Equal(t, instance.ID, "fwp-j3654")
}

func TestRemoveFirewallPolicy(t *testing.T) {
	instance := testLink[FirewallPolicy, FirewallPolicyGroup](
		t,
		(*Client).RemoveFirewallPolicy,
		"fwp-j3654",
		FirewallPolicyGroup{"grp-12345"},
		"firewall_policy",
		"POST",
		path.Join("firewall_policies", "fwp-j3654", "remove"),
		`{"server_group":"grp-12345"}`,
	)
	assert.Equal(t, instance.ID, "fwp-j3654")
}

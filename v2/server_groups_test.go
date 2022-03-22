package brightbox

import (
	"context"
	"path"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddServersToServerGroup(t *testing.T) {
	instance := testLink[ServerGroup, ServerGroupMemberList](
		t,
		(*Client).AddServersToServerGroup,
		"grp-sda44",
		ServerGroupMemberList{[]ServerGroupMember{{"srv-lv426"}}},
		"server_group",
		"POST",
		path.Join("server_groups", "grp-sda44", "add_servers"),
		`{"servers":[{"server":"srv-lv426"}]}`,
	)
	assert.Equal(t, instance.ID, "grp-sda44")
}

func TestRemoveServersFromServerGroup(t *testing.T) {
	instance := testLink[ServerGroup, ServerGroupMemberList](
		t,
		(*Client).RemoveServersFromServerGroup,
		"grp-sda44",
		ServerGroupMemberList{[]ServerGroupMember{{"srv-lv426"}}},
		"server_group",
		"POST",
		path.Join("server_groups", "grp-sda44", "remove_servers"),
		`{"servers":[{"server":"srv-lv426"}]}`,
	)
	assert.Equal(t, instance.ID, "grp-sda44")
}

func TestMoveServersToServerGroup(t *testing.T) {
	modify := (*Client).MoveServersToServerGroup
	from := "grp-sda44"
	to := "grp-12345"
	contents := ServerGroupMemberList{[]ServerGroupMember{{"srv-lv426"}}}
	jsonPath := "server_group"
	verb := "POST"
	expectedPath := path.Join("server_groups", "grp-sda44", "move_servers")
	expectedBody := `{"servers":[{"server":"srv-lv426"}],"destination":"grp-12345"}`

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
	instance, err := modify(client, context.Background(), from, to, contents)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, instance != nil)
	assert.Equal(t, instance.ID, "grp-sda44")
}

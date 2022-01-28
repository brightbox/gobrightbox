package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestUsers(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/users",
			ExpectBody:   "",
			GiveBody:     readJSON("users"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	p, err := All[User](client)
	assert.Assert(t, is.Nil(err), "All[User]() returned an error")
	assert.Assert(t, p != nil, "All[User]() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of users returned")
	ac := p[0]
	assert.Equal(t, "usr-kl435", ac.ID, "user id incorrect")
}

func TestUser(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/users/usr-kl435",
			ExpectBody:   "",
			GiveBody:     readJSON("user"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	ac, err := Instance[User](client, "usr-kl435")
	assert.Assert(t, is.Nil(err), "Instance[User] returned an error")
	assert.Assert(t, ac != nil, "Instance[User] returned nil")
	assert.Equal(t, "usr-kl435", ac.ID, "user id incorrect")
	assert.Equal(t, "John Jarvis", ac.Name, "user name incorrect")
}

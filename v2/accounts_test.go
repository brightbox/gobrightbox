package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestAccounts(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/accounts",
			ExpectBody:   "",
			GiveBody:     readJSON("accounts"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	p, err := All[Account](client)
	assert.Assert(t, is.Nil(err), "All[Account]() returned an error")
	assert.Assert(t, p != nil, "All[Account]() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of accounts returned")
	ac := p[0]
	assert.Equal(t, "acc-43ks4", ac.ID, "account id incorrect")
}

func TestAccount(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/accounts/acc-43ks4",
			ExpectBody:   "",
			GiveBody:     readJSON("account"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	ac, err := Instance[Account](client, "acc-43ks4")
	assert.Assert(t, is.Nil(err), "Instance[Account] returned an error")
	assert.Assert(t, ac != nil, "Instance[Account] returned nil")
	assert.Equal(t, "acc-43ks4", ac.ID, "account id incorrect")
	assert.Equal(t, "Brightbox", ac.Name, "account name incorrect")
}

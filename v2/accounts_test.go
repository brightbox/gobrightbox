package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestAccounts(t *testing.T) {
	instance := testAll[Account](
		t,
		"Account",
		"accounts",
		"account",
	)
	assert.Equal(t, instance.ID, "acc-43ks4")
}

func TestAccount(t *testing.T) {
	instance := testInstance[Account](
		t,
		"Account",
		"accounts",
		"account",
		"acc-43ks4",
	)
	assert.Equal(t, instance.Name, "Brightbox")
	assert.Equal(t, instance.Owner.ID, "usr-kl435")
}

func TestUpdateAccount(t *testing.T) {
	name := "Brightbox"
	uac := AccountOptions{ID: "acc-43ks4", Name: &name}
	_ = testUpdate[Account](
		t,
		"Account",
		"accounts",
		"account",
		"acc-43ks4",
		&uac,
		`{"name":"Brightbox"}`,
	)
}

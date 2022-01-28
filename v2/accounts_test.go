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
	assert.Equal(t, "acc-43ks4", instance.ID, "account id incorrect")
}

func TestAccount(t *testing.T) {
	instance := testInstance[Account](
		t,
		"Account",
		"accounts",
		"account",
		"acc-43ks4",
	)
	assert.Equal(t, "acc-43ks4", instance.ID, "account id incorrect")
	assert.Equal(t, "Brightbox", instance.Name, "account name incorrect")
}

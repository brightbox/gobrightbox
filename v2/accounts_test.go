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
}

package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestUsers(t *testing.T) {
	instance := testAll[User](
		t,
		"User",
		"users",
		"user",
	)
	assert.Equal(t, instance.ID, "usr-kl435")
}

func TestUser(t *testing.T) {
	instance := testInstance[User](
		t,
		"User",
		"users",
		"user",
		"usr-kl435",
	)
	assert.Equal(t, instance.ID, "usr-kl435")
	assert.Equal(t, instance.Name, "John Jarvis")
}

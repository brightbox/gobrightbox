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
	assert.Equal(t, "usr-kl435", instance.ID, "user id incorrect")
}

func TestUser(t *testing.T) {
	instance := testInstance[User](
		t,
		"User",
		"users",
		"user",
		"usr-kl435",
	)
	assert.Equal(t, "usr-kl435", instance.ID, "user id incorrect")
	assert.Equal(t, "John Jarvis", instance.Name, "user name incorrect")
}

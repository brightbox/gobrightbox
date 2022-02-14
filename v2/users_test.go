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
	assert.Equal(t, instance.Name, "John Jarvis")
	assert.Equal(t, instance.DefaultAccount.ID, "acc-43ks4")
}

func TestUpdateUser(t *testing.T) {
	name := "John Jarvis"
	uac := UserOptions{ID: "usr-kl435", Name: &name}
	_ = testUpdate[User](
		t,
		"User",
		"users",
		"user",
		&uac,
		`{"name":"John Jarvis"}`,
	)
}

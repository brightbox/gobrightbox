package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestCollaborations(t *testing.T) {
	instance := testAll[Collaboration](
		t,
		"Collaboration",
		"collaborations",
		"api client",
	)
	assert.Equal(t, instance.ID, "col-klek3")
}

func TestCollaboration(t *testing.T) {
	instance := testInstance[Collaboration](
		t,
		"Collaboration",
		"collaborations",
		"collaboration",
		"col-klek3",
	)
	assert.Equal(t, instance.Role, "admin")
	assert.Equal(t, instance.Account.ID, "acc-43ks4")
}

func TestCreateCollaboration(t *testing.T) {
	newAC := CollaborationOptions{}
	_ = testCreate[Collaboration](
		t,
		"Collaboration",
		"collaborations",
		"collaboration",
		"col-klek3",
		&newAC,
		"{}",
	)
}

func TestCreateCollaborationWithPermissionsGroup(t *testing.T) {
	email := "jason.null@example.com"
	newAC := CollaborationOptions{Email: &email}
	instance := testCreate[Collaboration](
		t,
		"Collaboration",
		"collaborations",
		"collaboration",
		"col-klek3",
		&newAC,
		`{"email":"jason.null@example.com"}`,
	)
	assert.Equal(t, instance.Email, "tclock@example.com")
}

func TestDestroyCollaboration(t *testing.T) {
	testDestroy[Collaboration](
		t,
		"Collaboration",
		"collaborations",
		"col-klek3",
	)
}

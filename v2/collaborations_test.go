package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateCollaborationWithPermissionsGroup(t *testing.T) {
	email := "jason.null@example.com"
	newResource := CollaborationOptions{Email: &email}
	instance := testModify(
		t,
		(*Client).CreateCollaboration,
		newResource,
		"collaboration",
		"POST",
		"collaborations",
		`{"email":"jason.null@example.com"}`,
	)
	assert.Equal(t, instance.Email, "tclock@example.com")
}

func TestResendCollaboration(t *testing.T) {
	instance := testModify(
		t,
		(*Client).ResendCollaboration,
		"col-klek3",
		"collaboration",
		"POST",
		path.Join("collaborations", "col-klek3", "resend"),
		"",
	)
	assert.Equal(t, instance.ID, "col-klek3")
}

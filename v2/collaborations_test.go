package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateCollaborationWithPermissionsGroup(t *testing.T) {
	email := "jason.null@example.com"
	newResource := CollaborationOptions{Email: &email}
	instance := testModify[Collaboration, *CollaborationOptions](
		t,
		(*Client).CreateCollaboration,
		&newResource,
		"collaboration",
		"POST",
		"collaborations",
		`{"email":"jason.null@example.com"}`,
	)
	assert.Equal(t, instance.Email, "tclock@example.com")
}

func TestResendCollaboration(t *testing.T) {
	testCommand(
		t,
		(*Client).ResendCollaboration,
		"col-klek3",
		"POST",
		path.Join("collaborations", "col-klek3", "resend"),
	)
}

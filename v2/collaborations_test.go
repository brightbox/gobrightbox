package brightbox

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateCollaborationWithPermissionsGroup(t *testing.T) {
	email := "jason.null@example.com"
	newResource := CollaborationOptions{Email: &email}
	instance := testModify[Collaboration, CollaborationOptions](
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

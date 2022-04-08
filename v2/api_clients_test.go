package brightbox

import (
	"testing"

	"github.com/brightbox/gobrightbox/v2/status/permissionsgroup"
	"gotest.tools/v3/assert"
)

func TestCreateAPIClientWithPermissionsGroup(t *testing.T) {
	pg := permissionsgroup.Full
	newResource := APIClientOptions{PermissionsGroup: pg}
	instance := testModify[APIClient, APIClientOptions](
		t,
		(*Client).CreateAPIClient,
		newResource,
		"api_client",
		"POST",
		"api_clients",
		`{"permissions_group":"full"}`,
	)
	assert.Equal(t, instance.PermissionsGroup, pg)
}

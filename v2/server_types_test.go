package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestServerTypes(t *testing.T) {
	instance := testAll[ServerType](
		t,
		"ServerType",
		"server_types",
		"server type",
	)
	assert.Equal(t, "typ-zx45f", instance.ID, "server_type id incorrect")
}

func TestServerType(t *testing.T) {
	instance := testInstance[ServerType](
		t,
		"ServerType",
		"server_types",
		"server_type",
		"typ-zx45f",
	)
	assert.Equal(t, "typ-zx45f", instance.ID, "server_type id incorrect")
	assert.Equal(t, "Small", instance.Name, "server_type name incorrect")
}

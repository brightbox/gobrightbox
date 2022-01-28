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
	assert.Equal(t, instance.ID, "typ-zx45f")
}

func TestServerType(t *testing.T) {
	instance := testInstance[ServerType](
		t,
		"ServerType",
		"server_types",
		"server_type",
		"typ-zx45f",
	)
	assert.Equal(t, instance.ID, "typ-zx45f")
	assert.Equal(t, instance.Name, "Small")
}

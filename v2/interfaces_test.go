package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestInterfaces(t *testing.T) {
	instance := testAll[Interface](
		t,
		"Interface",
		"interfaces",
		"interface",
	)
	assert.Equal(t, instance.ID, "int-ds42k")
}

func TestInterface(t *testing.T) {
	instance := testInstance[Interface](
		t,
		"Interface",
		"interfaces",
		"interface",
		"int-ds42k",
	)
	assert.Equal(t, instance.MacAddress, "02:24:19:00:00:ee")
}

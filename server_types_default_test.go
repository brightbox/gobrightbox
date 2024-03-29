// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestServerTypes(t *testing.T) {
	instance := testAll(
		t,
		(*Client).ServerTypes,
		"ServerType",
		"server_types",
		"ServerTypes",
	)
	assert.Equal(t, instance.ID, "typ-zx45f")
}

func TestServerType(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).ServerType,
		"ServerType",
		path.Join("server_types", "typ-zx45f"),
		"server_type",
		"typ-zx45f",
	)
	assert.Equal(t, instance.ID, "typ-zx45f")
}

func TestServerTypeByHandle(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).ServerType,
		"ServerType",
		path.Join("server_types", "typ-zx45f"),
		"server_type",
		"typ-zx45f",
	)
	assert.Equal(t, instance.ID, "typ-zx45f")
	handleInstance := testInstance[ServerType](
		t,
		(*Client).ServerTypeByHandle,
		"ServerType",
		path.Join("server_types"),
		"server_types",
		instance.Handle,
	)
	assert.Equal(t, instance.ID, handleInstance.ID)
}

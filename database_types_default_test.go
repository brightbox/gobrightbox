// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDatabaseServerTypes(t *testing.T) {
	instance := testAll(
		t,
		(*Client).DatabaseServerTypes,
		"DatabaseServerType",
		"database_types",
		"DatabaseServerTypes",
	)
	assert.Equal(t, instance.ID, "dbt-12345")
}

func TestDatabaseServerType(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).DatabaseServerType,
		"DatabaseServerType",
		path.Join("database_types", "dbt-12345"),
		"database_type",
		"dbt-12345",
	)
	assert.Equal(t, instance.ID, "dbt-12345")
}
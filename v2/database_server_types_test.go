package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestDatabaseServerTypes(t *testing.T) {
	instance := testAll[DatabaseServerType](
		t,
		"DatabaseServerType",
		"database_types",
		"database_type",
	)
	assert.Equal(t, instance.ID, "dbt-12345")
}

func TestDatabaseServerType(t *testing.T) {
	instance := testInstance[DatabaseServerType](
		t,
		"DatabaseServerType",
		"database_types",
		"database_type",
		"dbt-12345",
	)
	assert.Equal(t, instance.Default, false)
	assert.Equal(t, instance.Name, "Small")
}

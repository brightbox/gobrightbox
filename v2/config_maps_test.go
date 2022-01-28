package brightbox

import (
	"testing"

	"gotest.tools/assert"
)

func TestConfigMaps(t *testing.T) {
	instance := testAll[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"config map",
	)
	assert.Equal(t, instance.ID, "cfg-dsse2")
}

func TestConfigMap(t *testing.T) {
	instance := testInstance[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"config_map",
		"cfg-dsse2",
	)
	assert.Equal(t, instance.ID, "cfg-dsse2")
	assert.Equal(t, instance.Name, "example.test")
}

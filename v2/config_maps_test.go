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
	assert.Equal(t, "cfg-dsse2", instance.ID, "config map if incorrect")
}

func TestConfigMap(t *testing.T) {
	instance := testInstance[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"config_map",
		"cfg-dsse2",
	)
	assert.Equal(t, "cfg-dsse2", instance.ID, "config map if incorrect")
	assert.Equal(t, "example.test", instance.Name, "config map name incorrect")
}

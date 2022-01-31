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
	assert.Equal(t, instance.Name, "example.test")
}

func TestCreateConfigMap(t *testing.T) {
	newAC := ConfigMapOptions{}
	_ = testCreate[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"config_map",
		"cfg-dsse2",
		&newAC,
		"{}",
	)
}

func TestUpdateConfigMap(t *testing.T) {
	name := "dev client"
	uac := ConfigMapOptions{ID: "cfg-dsse2", Name: &name}
	_ = testUpdate[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"config_map",
		"cfg-dsse2",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyConfigMap(t *testing.T) {
	testDestroy[ConfigMap](
		t,
		"ConfigMap",
		"config_maps",
		"cfg-dsse2",
	)
}

package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConfigMaps(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/config_maps",
		ExpectBody:   "",
		GiveBody:     readJSON("config_maps"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	p, err := client.ConfigMaps()
	assert.NilError(t, err, "ConfigMaps() returned an error")
	assert.Assert(t, p != nil, "ConfigMaps() returned nil")
	assert.Equal(t, 1, len(p), "wrong number of config map returned")
	ac := p[0]
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID), "config map id incorrect")
}

func TestConfigMap(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "GET",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   "",
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	ac, err := client.ConfigMap("cfg-dsse2")
	assert.NilError(t, err, "ConfigMap() returned an error")
	assert.Assert(t, ac != nil, "ConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID), "config map id incorrect")
	assert.Check(t, is.Equal("example.test", ac.Name), "config map name incorrect")
}

func TestCreateConfigMap(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/config_maps",
		ExpectBody:   `{}`,
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	newAC := brightbox.ConfigMapOptions{}
	ac, err := client.CreateConfigMap(&newAC)
	assert.NilError(t, err, "CreateConfigMap() returned an error")
	assert.Assert(t, ac != nil, "CreateConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID))
}

func TestUpdateConfigMap(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   `{"name":"Hello","data":{"first":1,"second":"two","three":{"nest1":"one","nest2":2,"nest3":"maybe three"}}}`,
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := "Hello"
	data := map[string]interface{}{
		"first":  1,
		"second": "two",
		"three": map[string]interface{}{
			"nest1": "one",
			"nest2": 2,
			"nest3": "maybe three",
		},
	}
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Name: &name,
		Data: &data,
	}
	ac, err := client.UpdateConfigMap(&uac)
	assert.NilError(t, err, "UpdateConfigMap() returned an error")
	assert.Assert(t, ac != nil, "UpdateConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID))
}

func TestUpdateConfigMapToEmpty(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   `{"name":"","data":{}}`,
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := ""
	data := map[string]interface{}{}
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Name: &name,
		Data: &data,
	}
	ac, err := client.UpdateConfigMap(&uac)
	assert.NilError(t, err, "UpdateConfigMap() returned an error")
	assert.Assert(t, ac != nil, "UpdateConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID))
}

func TestUpdateConfigMapClearName(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   `{"name":""}`,
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	name := ""
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Name: &name,
	}
	ac, err := client.UpdateConfigMap(&uac)
	assert.NilError(t, err, "UpdateConfigMap() returned an error")
	assert.Assert(t, ac != nil, "UpdateConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID))
}

func TestUpdateConfigMapClearData(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "PUT",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   `{"data":{}}`,
		GiveBody:     readJSON("config_map"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	data := map[string]interface{}{}
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Data: &data,
	}
	ac, err := client.UpdateConfigMap(&uac)
	assert.NilError(t, err, "UpdateConfigMap() returned an error")
	assert.Assert(t, ac != nil, "UpdateConfigMap() returned nil")
	assert.Check(t, is.Equal("cfg-dsse2", ac.ID))
}

func TestDestroyConfigMap(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "DELETE",
		ExpectURL:    "/1.0/config_maps/cfg-dsse2",
		ExpectBody:   "",
		GiveBody:     "",
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err, "NewClient returned an error")

	err = client.DestroyConfigMap("cfg-dsse2")
	assert.NilError(t, err, "DestroyConfigMap() returned an error")
}

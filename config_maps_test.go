package gobrightbox_test

import (
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Nil(t, err, "NewClient returned an error")

	p, err := client.ConfigMaps()
	require.Nil(t, err, "ConfigMaps() returned an error")
	require.NotNil(t, p, "ConfigMaps() returned nil")
	require.Equal(t, 1, len(p), "wrong number of config map returned")
	ac := p[0]
	assert.Equal(t, "cfg-dsse2", ac.ID, "config map id incorrect")
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
	require.Nil(t, err, "NewClient returned an error")

	ac, err := client.ConfigMap("cfg-dsse2")
	require.Nil(t, err, "ConfigMap() returned an error")
	require.NotNil(t, ac, "ConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID, "config map id incorrect")
	assert.Equal(t, "example.test", ac.Name, "config map name incorrect")
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
	require.Nil(t, err, "NewClient returned an error")

	newAC := brightbox.ConfigMapOptions{}
	ac, err := client.CreateConfigMap(&newAC)
	require.Nil(t, err, "CreateConfigMap() returned an error")
	require.NotNil(t, ac, "CreateConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID)
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
	require.Nil(t, err, "NewClient returned an error")

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
	require.Nil(t, err, "UpdateConfigMap() returned an error")
	require.NotNil(t, ac, "UpdateConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID)
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
	require.Nil(t, err, "NewClient returned an error")

	name := ""
	data := map[string]interface{}{}
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Name: &name,
		Data: &data,
	}
	ac, err := client.UpdateConfigMap(&uac)
	require.Nil(t, err, "UpdateConfigMap() returned an error")
	require.NotNil(t, ac, "UpdateConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID)
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
	require.Nil(t, err, "NewClient returned an error")

	name := ""
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Name: &name,
	}
	ac, err := client.UpdateConfigMap(&uac)
	require.Nil(t, err, "UpdateConfigMap() returned an error")
	require.NotNil(t, ac, "UpdateConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID)
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
	require.Nil(t, err, "NewClient returned an error")

	data := map[string]interface{}{}
	uac := brightbox.ConfigMapOptions{
		ID:   "cfg-dsse2",
		Data: &data,
	}
	ac, err := client.UpdateConfigMap(&uac)
	require.Nil(t, err, "UpdateConfigMap() returned an error")
	require.NotNil(t, ac, "UpdateConfigMap() returned nil")
	assert.Equal(t, "cfg-dsse2", ac.ID)
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
	require.Nil(t, err, "NewClient returned an error")

	err = client.DestroyConfigMap("cfg-dsse2")
	require.Nil(t, err, "DestroyConfigMap() returned an error")
}

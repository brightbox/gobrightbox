package gobrightbox_test

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestError(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   `{"name":"new client"}`,
		GiveStatus:   403,
		GiveBody:     readJSON("api_error_forbidden"),
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()
	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err)

	name := "new client"
	newOptions := brightbox.APIClientOptions{Name: &name}
	instance, err := client.CreateAPIClient(&newOptions)
	assert.Assert(t, is.Nil(instance))
	var apierror *brightbox.APIError
	// apierror := new(brightbox.APIError)
	assert.ErrorType(t, err, apierror)
	if errors.As(err, &apierror) {
		assert.Equal(t, apierror.StatusCode, 403)
		assert.Equal(t, apierror.ErrorName, "Forbidden")
		assert.DeepEqual(
			t, apierror.Errors,
			[]string{"Account limit reached, please contact support for more information"},
		)
	}
}

func TestParseError(t *testing.T) {
	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/1.0/api_clients",
		ExpectBody:   `{"name":"new client"}`,
		GiveStatus:   200,
		GiveBody:     `{"foo": 1,}`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()
	client, err := brightbox.NewClient(ts.URL, "", nil)
	assert.NilError(t, err)

	name := "new client"
	newOptions := brightbox.APIClientOptions{Name: &name}
	instance, err := client.CreateAPIClient(&newOptions)
	assert.Assert(t, is.Nil(instance))
	jsonError := new(json.SyntaxError)
	if errors.As(err, &jsonError) {
		assert.Error(t, jsonError, "invalid character '}' looking for beginning of object key string")
	} else {
		assert.ErrorType(t, err, jsonError)
	}
}

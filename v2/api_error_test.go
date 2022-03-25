package brightbox

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestError(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   `{"name":"new client"}`,
			GiveStatus:   403,
			GiveBody:     readJSON("api_error_forbidden"),
		},
	)
	defer ts.Close()
	assert.NilError(t, err)

	name := "new client"
	instance, err := client.CreateAPIClient(context.Background(), APIClientOptions{Name: &name})
	assert.Assert(t, is.Nil(instance))
	apierror := new(APIError)
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
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   `{"name":"new client"}`,
			GiveStatus:   200,
			GiveBody:     `{"foo": 1,}`,
		},
	)
	defer ts.Close()
	assert.NilError(t, err)

	name := "new client"
	instance, err := client.CreateAPIClient(context.Background(), APIClientOptions{Name: &name})
	assert.Assert(t, is.Nil(instance))
	jsonError := new(json.SyntaxError)
	if errors.As(err, &jsonError) {
		assert.Error(t, jsonError, "invalid character '}' looking for beginning of object key string")
	} else {
		assert.ErrorType(t, err, jsonError)
	}
}

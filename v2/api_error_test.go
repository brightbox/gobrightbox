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
	assert.Error(t, err, "Forbidden: Account limit reached, please contact support for more information")
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

func TestAuthError(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   `{"name":"new client"}`,
			GiveStatus:   401,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.NilError(t, err)

	name := "new client"
	instance, err := client.CreateAPIClient(context.Background(), APIClientOptions{Name: &name})
	assert.ErrorContains(t, err, ": 401 Unauthorized")
	assert.Assert(t, is.Nil(instance))
	apierror := new(APIError)
	assert.ErrorType(t, err, apierror)
	if errors.As(err, &apierror) {
		assert.Equal(t, apierror.StatusCode, 401)
		assert.Equal(t, apierror.ErrorName, "")
		assert.Assert(t, is.Len(apierror.Errors, 0))
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
	assert.ErrorContains(t, err, ": invalid character '}' looking for beginning of object key string")
	assert.Assert(t, is.Nil(instance))
	jsonError := new(json.SyntaxError)
	if errors.As(err, &jsonError) {
		assert.Error(t, jsonError, "invalid character '}' looking for beginning of object key string")
	} else {
		assert.ErrorType(t, err, jsonError)
	}
}

func TestParseShortfall(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/api_clients",
			ExpectBody:   `{"name":"new client"}`,
			GiveStatus:   200,
			GiveBody:     `{"id": "1"}{"id": "2"}`,
		},
	)
	defer ts.Close()
	assert.NilError(t, err)

	name := "new client"
	instance, err := client.CreateAPIClient(context.Background(), APIClientOptions{Name: &name})
	assert.ErrorContains(t, err, ": Response body has additional unparsed data at position 12")
	assert.Assert(t, is.Nil(instance))
}

func TestUnmarshalError(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "GET",
			ExpectURL:    "/1.0/servers/srv-testy",
			ExpectBody:   "",
			GiveStatus:   200,
			GiveBody:     `{"status": "available"}`,
		},
	)
	defer ts.Close()
	assert.NilError(t, err)

	instance, err := client.Server(context.Background(), "srv-testy")
	assert.ErrorContains(t, err, ": json: cannot unmarshal available into Go struct field Server.status of type serverstatus.Enum")
	assert.Assert(t, is.Nil(instance))
	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		assert.Equal(t, unmarshalError.Offset, int64(23))
		assert.Error(t, unmarshalError, "json: cannot unmarshal available into Go struct field Server.status of type serverstatus.Enum")
	}
}

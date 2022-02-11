package gobrightbox_test

import (
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	brightbox "github.com/brightbox/gobrightbox"
)

func TestGET(t *testing.T) {

	handler := APIMock{T: t, ExpectMethod: "GET", ExpectURL: "/some/path"}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.MakeAPIRequest("GET", "/some/path", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("returned %+v, want %+v",
			res.StatusCode, 200)
	}

}

func TestPOST(t *testing.T) {

	handler := APIMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectURL:    "/some/resource",
		ExpectBody:   `{"hello":"world"}`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.MakeAPIRequest("POST", "/some/resource", map[string]string{"hello": "world"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAPIError403(t *testing.T) {
	handler := APIMock{
		GiveStatus: 403,
		GiveBody: `{
"error":"error title",
"error_description": "error desc",
"error_name": "name",
"errors": ["one", "two"],
"another_key": "ignored"}`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.MakeAPIRequest("POST", "/some/resource", nil, nil)
	if err == nil {
		t.Fatal("Expected error")
	}
	errt := reflect.TypeOf(err).String()
	if errt != "*gobrightbox.APIError" {
		t.Fatalf("Returned error was %q, wanted *gobrightbox.APIError", errt)
	}
	var apierr *brightbox.APIError
	if errors.As(err, &apierr) {
		u := apierr.RequestURL.RequestURI()
		if u != "/some/resource" {
			t.Fatalf("err.RequestURL was %q, wanted /some/resource", u)
		}
		if apierr.StatusCode != 403 {
			t.Fatalf("err.StatusCode was %d, wanted 403", apierr.StatusCode)
		}
		if apierr.AuthError != "error title" {
			t.Fatalf("err.AuthError was %q", apierr.AuthError)
		}
		if apierr.AuthErrorDescription != "error desc" {
			t.Fatalf("err.AuthErrorDescription was %q", apierr.AuthError)
		}
		if apierr.ErrorName != "name" {
			t.Fatalf("err.ErrorName was %q", apierr.ErrorName)
		}
		if len(apierr.Errors) != 2 {
			t.Fatalf("err.Errors was %q", apierr.Errors)
		}
	}
}

func TestAPIError500WithoutErrorJson(t *testing.T) {
	handler := APIMock{
		GiveStatus: 500,
		GiveBody:   `some error message`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.MakeAPIRequest("POST", "/some/resource", nil, nil)
	if err == nil {
		t.Fatal("Expected error")
	}
	errt := reflect.TypeOf(err).String()
	if errt != "*gobrightbox.APIError" {
		t.Fatalf("Returned error was %q, wanted *gobrightbox.APIError", errt)
	}
	var apierr *brightbox.APIError
	if errors.As(err, &apierr) {
		u := apierr.RequestURL.RequestURI()
		if u != "/some/resource" {
			t.Fatalf("err.RequestURL was %q, wanted /some/resource", u)
		}
	}
}

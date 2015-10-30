package brightbox_test

import (
	"github.com/brightbox/gobrightbox"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGET(t *testing.T) {

	handler := ApiMock{T: t, ExpectMethod: "GET", ExpectUrl: "/some/path"}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.MakeApiRequest("GET", "/some/path", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("returned %+v, want %+v",
			res.StatusCode, 200)
	}

}

func TestPOST(t *testing.T) {

	handler := ApiMock{
		T:            t,
		ExpectMethod: "POST",
		ExpectUrl:    "/some/resource",
		ExpectBody:   `{"hello":"world"}`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.MakeApiRequest("POST", "/some/resource", map[string]string{"hello": "world"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApiError403(t *testing.T) {
	handler := ApiMock{
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

	_, err = client.MakeApiRequest("POST", "/some/resource", nil, nil)
	if err == nil {
		t.Fatal("Expected error")
	}
	errt := reflect.TypeOf(err).String()
	if errt != "brightbox.ApiError" {
		t.Fatalf("Returned error was %q, wanted ApiError", errt)
	}
	apierr := err.(brightbox.ApiError)
	u := apierr.RequestUrl.RequestURI()
	if u != "/some/resource" {
		t.Fatalf("err.RequestUrl was %q, wanted /some/resource", u)
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

func TestApiError500WithoutErrorJson(t *testing.T) {
	handler := ApiMock{
		GiveStatus: 500,
		GiveBody:   `some error message`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.MakeApiRequest("POST", "/some/resource", nil, nil)
	if err == nil {
		t.Fatal("Expected error")
	}
	errt := reflect.TypeOf(err).String()
	if errt != "brightbox.ApiError" {
		t.Fatalf("Returned error was %q, wanted ApiError", errt)
	}
	apierr := err.(brightbox.ApiError)
	u := apierr.RequestUrl.RequestURI()
	if u != "/some/resource" {
		t.Fatalf("err.RequestUrl was %q, wanted /some/resource", u)
	}
}

func TestParseError(t *testing.T) {
	handler := ApiMock{
		GiveStatus: 200,
		GiveBody:   `{"name": 1000}`,
	}
	ts := httptest.NewServer(&handler)
	defer ts.Close()

	client, err := brightbox.NewClient(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	o := make(map[string]string)
	_, err = client.MakeApiRequest("GET", "/some/resource", nil, &o)
	if err == nil {
		t.Fatal("Expected error")
	}
	errt := reflect.TypeOf(err).String()
	if errt != "brightbox.ApiError" {
		t.Fatalf("Returned error was %q, wanted ApiError", errt)
	}
	apierr := err.(brightbox.ApiError)
	if apierr.ParseError == nil {
		t.Fatalf("err.ParseError was nil")
	}
}

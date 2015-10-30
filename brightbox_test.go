package brightbox_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func readJson(name string) string {
	content, err := ioutil.ReadFile("testdata/" + name + ".json")
	if err != nil {
		panic(err)
	}
	return string(content)
}

type ApiMock struct {
	*testing.T
	ExpectMethod string
	ExpectUrl    string
	ExpectBody   string
	GiveStatus   int
	GiveBody     string
}

func (a *ApiMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.ExpectMethod != "" && r.Method != a.ExpectMethod {
		a.Fatalf("Expected method %q but got %q", a.ExpectMethod, r.Method)
	}
	if a.ExpectUrl != "" && r.URL.String() != a.ExpectUrl {
		a.Fatalf("Expected url %q but got %q", a.ExpectUrl, r.URL.String())
	}

	if a.ExpectBody != "" {
		b, _ := ioutil.ReadAll(r.Body)
		tb := strings.TrimSpace(string(b))
		if a.ExpectBody != tb {
			a.Fatalf("Expected request body %q but got %q", a.ExpectBody, tb)
		}
	}

	if a.GiveStatus > 0 {
		w.WriteHeader(a.GiveStatus)
	} else {
		w.WriteHeader(200)
	}
	w.Header().Set("Content-Type", "application/json")

	if a.GiveBody != "" {
		w.Write([]byte(a.GiveBody))
	}
}

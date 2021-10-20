package gobrightbox_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readJSON(name string) string {
	content, err := ioutil.ReadFile("testdata/" + name + ".json")
	if err != nil {
		panic(err)
	}
	return string(content)
}

type APIMock struct {
	*testing.T
	ExpectMethod string
	ExpectURL    string
	ExpectBody   interface{}
	GiveStatus   int
	GiveBody     string
	GiveHeaders  map[string]string
}

func (a *APIMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.ExpectMethod != "" && r.Method != a.ExpectMethod {
		require.Equal(a, a.ExpectMethod, r.Method, "method didn't match")
	}
	if a.ExpectURL != "" && r.URL.String() != a.ExpectURL {
		require.Equal(a, a.ExpectURL, r.URL.String(), "url didn't match")
	}

	switch expectBody := a.ExpectBody.(type) {
	case string:
		b, _ := ioutil.ReadAll(r.Body)
		tb := strings.TrimSpace(string(b))
		if expectBody != tb {
			a.Fatalf("Expected request body %q but got %q", expectBody, tb)
		}
	case map[string]interface{}:
	case map[string]string:
		var decodedReqBody map[string]interface{}
		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(b, &decodedReqBody)
		if err != nil {
			a.Fatalf("Couldn't parse request body json: %s", err)
		}
		for key, value := range expectBody {
			decodedVal, ok := decodedReqBody[key]
			if !ok {
				a.Errorf("Expected key %q in request body but was missing", key)
			} else if fmt.Sprintf("%s", expectBody[key]) != fmt.Sprintf("%s", decodedVal) {
				assert.Equal(a, value, decodedReqBody[key], fmt.Sprintf("Key %q in request body doesn't match", key))
			}
		}
		for key, _ := range decodedReqBody {
			_, ok := expectBody[key]
			if !ok {
				a.Errorf("Unexpected key %q found in request body json", key)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	for k, v := range a.GiveHeaders {
		w.Header().Set(k, v)
	}

	if a.GiveStatus > 0 {
		w.WriteHeader(a.GiveStatus)
	} else {
		w.WriteHeader(200)
	}
	if a.GiveBody != "" {
		w.Write([]byte(a.GiveBody))
	}
}

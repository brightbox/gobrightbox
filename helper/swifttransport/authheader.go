package swifttransport

import (
	"net/http"

	"golang.org/x/oauth2"
)

// setAuthHeader sets the Swift Authorization header to r using the access
// token in t.
//
func setAuthHeader(t *oauth2.Token, r *http.Request) {
	r.Header.Set("X-Auth-Token", t.AccessToken)
}

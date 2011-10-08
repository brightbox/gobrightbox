/*
Copyright 2011 Brightbox Systems Ltd.

This file is part of Brightbox.go

brightbox.go is free software: you can redistribute it and/or modify
it under the terms of the Lesser GNU General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the Lesser
GNU General Public License for more details.
	
You should have received a copy of the Lesser GNU General Public
License along with this program.  If not, see
<http://www.gnu.org/licenses/>
*/

package brightbox

import (
	"os"
	"time"
	"http"
	"json"
	"strings"
	"io/ioutil"
)

// ApiClientAuth represents an ApiClient used for OAuth authentication
type ApiClientAuth struct {
	Id         string
	secret     string
	url        string
	token      string
	expires    int64
}

// NewApiClientAuth returns a new ApiClientAuth structure instantiated
// with the given data
func NewApiClientAuth(url string, id string, secret string) *ApiClientAuth {
	c := new(ApiClientAuth)
	c.url = url
	c.Id = id
	c.secret = secret
	return c
}

// Token() returns the OAuth token. If it has no token, or the token
// is expired, it requests one using RequestToken()
func (auth *ApiClientAuth) Token() (string, int64, os.Error) {
	if auth.token == "" || auth.expires < time.Seconds() + 60 {
		err := auth.RequestToken()
		if err != nil {
			return "", 0, err
		}
	}
	return auth.token, auth.expires, nil
}

func (auth *ApiClientAuth) SetToken(token string, expires int64) os.Error {
	if token == "" {
		return ErrInvalidToken
	}
	if expires < time.Seconds() + 60 {
		return ErrTokenExpired
	}
	auth.token = token
	auth.expires = expires
	return nil
}

func (auth *ApiClientAuth) String() string {
	return auth.Id
}

// ApiClientAuth.RequestToken
func (auth *ApiClientAuth) RequestToken() os.Error {
	token, expires, err := RequestToken(auth)
	if err != nil {
		return err
	}
	return auth.SetToken(token, expires)
}

// RequestToken makes an authenticated request to the API and returns
// the OAuth token and the expiry time
func RequestToken(auth *ApiClientAuth) (string, int64, os.Error) {
	treq := map[string]string{}
	treq["client_id"] = auth.Id
	treq["grant_type"] = "none"

	var s []uint8
	var err os.Error
	s, err = json.Marshal(treq)

	req, err := http.NewRequest("POST", auth.url + "/token", strings.NewReader(string(s)))
	defer req.Body.Close()
	req.SetBasicAuth(auth.Id, auth.secret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "brightbox.go")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", 0, err
	}

	s, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	var tokres map[string]interface{}
	err = json.Unmarshal(s, &tokres)
	if err != nil {
		return "", 0, err
	}

	token := tokres["access_token"]
	if token != nil && token != "" {
		// BUG(johnl): should use expiry time from server
		return token.(string), time.Seconds() + 7200, nil
	}
	return "", 0, os.NewError("Token not granted")
}

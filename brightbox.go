/*
Copyright 2011 Brightbox Systems Ltd.

This program is free software: you can redistribute it and/or modify
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
	"http"
	"json"
	"strings"
	"os"
	"io/ioutil"
)

type Authenticator interface {
	RequestToken() os.Error
	Token() (string, os.Error)
}

type ApiClientAuth struct {
	id         string
	secret     string
	url        string
	token      string
}

type Client struct {
	auth       Authenticator
	url        string
	version    string
}

func NewClient(url string, version string, auth Authenticator) *Client {
	c := new(Client)
	c.url = url
	c.auth = auth
	c.version = version
	return c
}

func NewApiClientAuth(url string, id string, secret string) *ApiClientAuth {
	c := new(ApiClientAuth)
	c.url = url
	c.id = id
	c.secret = secret
	return c
}

func (client *Client) DoRequest(method string, path string, body string) ([]interface{}, *http.Response, os.Error) { 
	var s []uint8
	var res *http.Response
	var err os.Error
	var token string
	var req *http.Request
	token, err = client.auth.Token()
	if err != nil {
		return nil, nil, err
	}
	req, err = http.NewRequest(method, client.url + "/" + client.version + path, strings.NewReader(body))
	req.Header.Set("Authorization", "OAuth " + token)
	res, err = http.DefaultClient.Do(req)
	// FIXME: If this errors due to expired token, we should get a new token and retry
	if err != nil {
		return nil, res, err
	}
	s, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}
	res.Body.Close()
	var rjson []interface{}

	err = json.Unmarshal(s, &rjson)
	if err != nil {
		return nil, res, err
	}
	return rjson, res, err
}

func (auth *ApiClientAuth) Token() (string, os.Error) {
	if auth.token == "" {
		err := auth.RequestToken()
		if err != nil {
			return "", err
		}
	}
	return auth.token, nil
}

func (auth *ApiClientAuth) RequestToken() os.Error {
	token, err := RequestToken(auth)
	if err != nil {
		return err
	}
	auth.token = token
	return nil
}

func RequestToken(auth *ApiClientAuth) (string, os.Error) {
	treq := map[string]string{}
	treq["client_id"] = auth.id
	treq["grant_type"] = "none"

	var s []uint8
	var err os.Error
	s, err = json.Marshal(treq)

	req, err := http.NewRequest("POST", auth.url + "/token", strings.NewReader(string(s)))
	defer req.Body.Close()
	req.SetBasicAuth(auth.id, auth.secret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "brightbox.go")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	s, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var tokres map[string]interface{}
	err = json.Unmarshal(s, &tokres)
	if err != nil {
		return "", err
	}

	token := tokres["access_token"]
	if token != nil && token != "" {
		return token.(string), nil
	}
	return "", os.NewError("Token not granted")
}
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

// The brightbox package provides a GO interface to the Brightbox
// Cloud API. See http://brightbox.com or
// http://docs.brightbox.com/reference/api/ for more details.
package brightbox

import (
	"http"
	"json"
	"strings"
	"os"
	"io/ioutil"
	"time"
	"fmt"
	"compress/gzip"
)


var (
  ErrTokenExpired      = os.NewError("brightbox: authentication token has expired")
  ErrInvalidToken      = os.NewError("brightbox: invalid token")
)

type Authenticator interface {
	RequestToken() os.Error
	Token() (string, int64, os.Error)
	SetToken(string, int64) os.Error
	String()   string
}

// ApiClientAuth represents an ApiClient used for OAuth authentication
type ApiClientAuth struct {
	Id         string
	secret     string
	url        string
	token      string
	expires    int64
}

// Client represents a connection to the API with a given Authenticator
type Client struct {
	auth       Authenticator
	url        string
	version    string
}

type ServerType struct {
	Id         string
	Handle     string
}

type Zone struct {
	Id         string
  Handle     string
}

// Server represents a Cloud Server
type Server struct {
	Id         string
	Status     string
	ServerType ServerType
	Zone       Zone
	// Image   Image
  Name       string
}

// NewClient returns a new Client structure instantiated with the
// given API url, the API version and Authenticator
func NewClient(url string, version string, auth Authenticator) *Client {
	c := new(Client)
	c.url = url
	c.auth = auth
	c.version = version
	return c
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

// DoRequest makes an API request with the given HTTP method ("GET",
// "POST" etc) to the given path ("/servers") and sends the given
// body. It returns parsed json, the original http response and an
// error object if there was a problem.
func (client *Client) DoRequest(method string, path string, body string) (interface{}, *http.Response, os.Error) { 
	var (
		s []uint8
		res *http.Response
		err os.Error
		token string
		req *http.Request
	)
	token, _, err = client.auth.Token()
	if err != nil {
		return nil, nil, err
	}
	req, err = http.NewRequest(method, client.url + "/" + client.version + path, strings.NewReader(body))
	req.Header.Set("Authorization", "OAuth " + token)
	req.Header.Set("Accept-Encoding", "gzip")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()
	if res.Header.Get("Content-Encoding") == "gzip" {
		var ungzipped_body *gzip.Decompressor
		ungzipped_body, err = gzip.NewReader(res.Body)
		defer ungzipped_body.Close()
		s, err = ioutil.ReadAll(ungzipped_body)
	} else	{
		s, err = ioutil.ReadAll(res.Body)
	}
	if err != nil {
		return nil, res, err
	}

	var rjson interface{}

	err = json.Unmarshal(s, &rjson)
	if err != nil {
		return nil, res, err
	}
	return rjson, res, err
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

func NewZoneFromJson(json_data interface{}) (Zone, os.Error) {
	var j_z = json_data.(map[string]interface{})
	z := new(Zone)
	z.Id = j_z["id"].(string)
	z.Handle = j_z["handle"].(string)
	return *z, nil
}

func NewServerTypeFromJson(json_data interface{}) (ServerType, os.Error) {
	var j_st = json_data.(map[string]interface{})
	st := new(ServerType)
	st.Id = j_st["id"].(string)
	if j_st["handle"] == nil {
		st.Handle = st.Id
	} else {
		st.Handle = j_st["handle"].(string)
	}
	// BUG(johnl): Other fields
	return *st, nil
}

// NewServerFromJson returns a new Server structure instantiated from
// json data returned by the API
func NewServerFromJson(json_data interface{}) (Server, os.Error) {
	var j_server = json_data.(map[string]interface{})
	var err os.Error
	server := new(Server)
	server.Id = j_server["id"].(string)
	server.Status = j_server["status"].(string)
	server.ServerType, err = NewServerTypeFromJson(j_server["server_type"])
	if err != nil {
		panic(err)
	}
	server.Zone, err = NewZoneFromJson(j_server["zone"])
	server.Name = j_server["name"].(string)
	return *server, nil
}

// ListServers gets a list of the Servers via the API and returns a
// slice of Server structures.
func (client *Client) ListServers() []Server {
	j_servers, _, err := client.DoRequest("GET", "/servers", "")
	if err != nil {
		panic(err)
	}
	servers := make([]Server, len(j_servers.([]interface{})))
	for i, s := range j_servers.([]interface{}) {
		server, err := NewServerFromJson(s)
		if err != nil {
			panic(err)
		}
		servers[i] = server
	}
	return servers
}

// SetupAuthenticatorCache tries to read a cached token from the local filesystem
func SetupAuthenticatorCache(auth Authenticator) os.Error {
	var (
		err       os.Error
		token     string
		expires   int64
		f         *os.File
	)
	cache_filename := "/home/john/.brightbox/" + auth.String() + ".oauth_token.v2"
	f, err = os.Open(cache_filename)
	if f != nil {
		_, err = fmt.Fscanf(f, "%s", &token)
		_, err = fmt.Fscanf(f, "%d", &expires)
		f.Close()
	}
	if auth.SetToken(token, expires) != nil {
		token, expires, err = auth.Token()
		if err != nil {
			return nil
		}
		// BUG(johnl): should write to temp file and rename
		ioutil.WriteFile(cache_filename, []uint8(fmt.Sprintf("%s %d", token, expires)), 0600)
	}
	return nil
}
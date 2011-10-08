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

type Image struct {
	Id         string
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

func NewImageFromJson(json_data interface{}) (Image, os.Error) {
	var j_img = json_data.(map[string]interface{})
	img := new(Image)
	img.Id = j_img["id"].(string)
	return *img, nil
}

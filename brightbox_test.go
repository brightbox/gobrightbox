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
	"testing"
	"os"
)

func TestNewApiClientAuth(t *testing.T) {
	c := NewApiClientAuth("auth", "cli-xxxxx", "asdf1234")
	if c.url != "auth" {
		t.Error("url not set correctly")
	}
	if c.id != "cli-xxxxx" {
		t.Error("id not set correcty")
	}
	if c.secret != "asdf1234" {
		t.Error("secret not set correctly")
	}
	if c.token != "" {
		t.Error("token was not default empty")
	}
}

func TestRequestTokenWithInvalidDetails(t *testing.T) {
	c := NewApiClientAuth("https://api.gb1.brightbox.com", "test", "asdf1234")
	token, err := RequestToken(c)
	if token != "" {
		t.Errorf("token should be empty")
	}
	if err == nil || err.String() != "Token not granted" {
		t.Errorf("err should be 'Token not granted'")
	}
}

func TestRequestToken(t *testing.T) {
	if os.Getenv("CLIENT") == "" || os.Getenv("SECRET") == "" {
		t.Fatal("Test requires CLIENT and SECRET env variables set")
	}
	c := NewApiClientAuth("https://api.gb1.brightbox.com", os.Getenv("CLIENT"), os.Getenv("SECRET"))
	token, err := RequestToken(c)
	if token == "" {
		t.Errorf("token should not be nil")
	}
	if err != nil {
		t.Errorf("err should be nil")
	}
}

func TestNewClient(t *testing.T) {
	auth := NewApiClientAuth("https://api.gb1.brightbox.com", os.Getenv("CLIENT"), os.Getenv("SECRET"))
	c := NewClient("https://api.gb1.brightbox.com", "1.0", auth)
	if c.auth != auth {
		t.Errorf("auth not set")
	}
	if c.url != "https://api.gb1.brightbox.com" {
		t.Errorf("url not set")
	}
}

func TestClientDoRequest(t *testing.T) {
	auth := NewApiClientAuth("https://api.gb1.brightbox.com", os.Getenv("CLIENT"), os.Getenv("SECRET"))
	c := NewClient("https://api.gb1.brightbox.com", "1.0", auth)
	body, res, err := c.DoRequest("GET", "/servers", "")
	if body == nil {
		t.Errorf("response shouldn't be empty")
	}
	if res == nil {
		t.Errorf("res shouldn't be nil")
	}
	if err != nil {
		t.Errorf("err should be nil: %s", err)
	}
}

func TestClientDoRequestInvalid(t *testing.T) {
	auth := NewApiClientAuth("https://api.gb1.brightbox.com", os.Getenv("CLIENT"), os.Getenv("SECRET"))
	c := NewClient("https://api.gb1.brightbox.com", "1.0", auth)
	body, res, err := c.DoRequest("GET", "/test", "")
	if body == nil {
		t.Errorf("response shouldn't be empty")
	} else if body.(map[string]interface{})["error_name"] == nil {
		t.Errorf("error response not parsed")
	}
	if res == nil {
		t.Errorf("res shouldn't be nil")
	} else if res.StatusCode != 404 {
		t.Errorf("res.StatusCode should be 404")
	}
	if err != nil {
		t.Errorf("err should be nil: %s", err)
	}
}

func TestClientDoRequestWithInvalidAuth(t *testing.T) {
	auth := NewApiClientAuth("https://api.gb1.brightbox.com", "cli-test", "brokensecret")
	c := NewClient("https://api.gb1.brightbox.com", "1.0", auth)
	body, res, err := c.DoRequest("GET", "/servers", "")
	if body != nil {
		t.Errorf("response should be empty")
	}
	if res != nil {
		t.Errorf("res shouldn't be nil")
	}
	if err == nil {
		t.Errorf("err should not be nil")
	}
}
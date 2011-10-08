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
	"os"
)

// Server represents a Cloud Server
type Server struct {
	Id         string
	Status     string
	ServerType ServerType
	Zone       Zone
	Image      Image
  Name       string
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
	server.Image, err = NewImageFromJson(j_server["image"])
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


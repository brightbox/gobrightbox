# Brightbox Golang Client

`gobrightbox` is a [Brightbox Cloud](https://www.brightbox.com) [API](https://api.gb1.brightbox.com/1.0/)
client implementation written in [Go](http://golang.org/).

Documentation is available at [godoc.org](http://godoc.org/github.com/brightbox/gobrightbox).

## Authentication

This client does not itself handle authentication. Instead, use the standard
[OAuth2](https://godoc.org/golang.org/x/oauth2) golang library to
[authenticate](https://api.gb1.brightbox.com/1.0/#authentication) and create
tokens.

## Currently implemented

* Full [Server](https://api.gb1.brightbox.com/1.0/#server) support
* Full [Server Group](https://api.gb1.brightbox.com/1.0/#server_group) support
* Full [CloudIP](https://api.gb1.brightbox.com/1.0/#cloud_ip) support
* Basic [Image](https://api.gb1.brightbox.com/1.0/#image) support
* Basic [Api Client](https://api.gb1.brightbox.com/1.0/#api_client) support
* Basic event stream support

## TODO

* Load Balancer support
* Firewall policy support
* Orbit storage support
* Collaboration support
* Cloud SQL support

## Help

If you need help using this library, drop an email to support at brightbox dot com.

## License

This code is released under an MIT License.

Copyright (c) 2015 Brightbox Systems Ltd.

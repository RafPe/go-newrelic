package core

import (
	resty "gopkg.in/resty.v1"
)

// A Config provides configuration to a service client instance.
type Config struct {
	Endpoint string
}

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	Config Config
}

// A APIRequest is used to package data/params and operation
// used by client to send the request
type APIRequest struct {
	Method string
}

//newHTTPrequest
func newHTTPrequest() {
	_, _ = resty.R().Get("http://httpbin.org/get")

	return
}

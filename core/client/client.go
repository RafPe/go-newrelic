package client

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
	Resty  *resty.Client
}

//New
func New(cfg Config) *Client {
	svc := &Client{
		Config: cfg,
		Resty:  resty.New(),
	}

	svc.Resty.SetDebug(true)

	return svc
}

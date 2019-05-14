package client

import (
	"fmt"

	u "github.com/RafPe/go-newrelic/core/nrq"
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

//newHTTPrequest
func newHTTPrequest(x u.Nrq) {
	fmt.Println(x)
	_, _ = resty.R().Get("http://httpbin.org/get")

	return
}

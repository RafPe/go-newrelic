package client

import (
	"fmt"
	"os"

	"github.com/RafPe/go-newrelic/core/requests"
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

//New TODO:document
func New(cfg Config) *Client {
	svc := &Client{
		Config: cfg,
		Resty:  resty.New(),
	}

	svc.Resty.SetDebug(true)

	return svc
}

//ExecuteRequest TODO:document
func (cl *Client) ExecuteRequest(newRequest *requests.Nrq) (*resty.Response, error) {

	req := cl.Resty.R().
		SetQueryParams(newRequest.QueryParams).
		SetHeaders(newRequest.Headers).
		SetHeader("X-Api-Key", os.Getenv("NEWRELIC_APIKEY")).
		SetResult(newRequest.ResponseModel)

	resp, err := req.Execute(resty.MethodGet, fmt.Sprintf("%s/%s", cl.Config.Endpoint, newRequest.URLPath))

	return resp, err
}

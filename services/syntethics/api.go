package syntehics

import (
	"fmt"
	"os"

	resty "gopkg.in/resty.v1"
)

// GetAllMonitors returns all monitors
// configured on a given account
func (c *Syntethics) GetAllMonitors() *SyntethicMonitors {

	req := c.Client.Resty.R().
		SetQueryParams(map[string]string{
			"limit":  "100",
			"offset": "0",
		}).
		SetHeader("X-Api-Key", os.Getenv("NEWRELIC_APIKEY")).
		SetResult(&SyntethicMonitors{})

	resp, err := req.Execute(resty.MethodGet, "https://synthetics.newrelic.com/synthetics/api/v3/monitors")

	resp.Header()

	fmt.Println(err, resp.Result())

	return nil
}

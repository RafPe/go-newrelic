package syntehics

import (
	"fmt"

	"github.com/RafPe/go-newrelic/core/requests"
)

// GetAllMonitors returns all monitors
// configured on a given account
func (c *Syntethics) GetAllMonitors() *SyntethicMonitors {

	newRequest := requests.Nrq{
		Method: "GET",
		QueryParams: map[string]string{
			"limit":  "100",
			"offset": "0",
		},
		RequestModel: &SyntethicMonitors{},
		Headers:      nil,
	}

	resp, err := c.Client.ExecuteRequest(&newRequest)

	fmt.Println(resp, err)

	return nil
}

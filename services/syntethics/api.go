package syntehics

import (
	"fmt"
	"strconv"

	"github.com/RafPe/go-newrelic/core/requests"
)

// GetAllMonitors returns all monitors
// configured on a given account
func (c *Syntethics) GetAllMonitors() *SyntethicMonitors {
	var offset int
	var limit = 3 // TODO: Move to const ? Global config var?
	var syntethicMonitorsArr []SyntethicMonitor

	newRequest := requests.Nrq{
		Method: "GET",
		QueryParams: map[string]string{
			"limit":  strconv.Itoa(limit),
			"offset": strconv.Itoa(offset),
		},
		ResponseModel: &SyntethicMonitors{},
		Headers:       nil,
	}

	// Executes first request to get all monitors.
	resp, err := c.Client.ExecuteRequest(&newRequest)
	if err != nil {
		panic(err)
	}

	// Process first response from getting all monitors
	syntethicMonitorsResp := resp.Result().(*SyntethicMonitors)
	for _, k := range syntethicMonitorsResp.Monitors {
		syntethicMonitorsArr = append(syntethicMonitorsArr, k)
	}

	// Calculate the total number of all monitors based
	// on the X-Total-Count header
	totalCountStr := resp.Header()["X-Total-Count"][0]
	totalCount, err := strconv.Atoi(totalCountStr)
	if err != nil {
		panic(err)
	}

	// if 119 > 100
	if totalCount > limit {
		offset += limit

		for offset < totalCount {
			fmt.Println("Calling again")

			// Set value to 101
			newRequest.QueryParams["offset"] = strconv.Itoa(offset)

			// Call with new values
			resp, err := c.Client.ExecuteRequest(&newRequest)
			if err != nil {
				panic(err)
			}
			fmt.Println(resp, err)

			// offset from 0 needs to go into 101
			// offset is now 101
			offset += limit
		}

	}

	return nil
}

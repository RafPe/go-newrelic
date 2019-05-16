package syntehics

import (
	"github.com/RafPe/go-newrelic/core/client"
)

// Syntethics provides the API operation methods for making requests to
// New Relic Syntethics. See this package's package overview docs
// for details on the service.
type Syntethics struct {
	*client.Client
}

// New creates a new instance of the Amplify client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Amplify client from just a session.
//     svc := amplify.New(mySession)
//
//     // Create a Amplify client with additional configuration
//     svc := amplify.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New() *Syntethics {
	return newClient()
}

// newClient creates, initializes and returns a new service client instance.
func newClient() *Syntethics {
	svc := &Syntethics{
		Client: client.New(client.Config{
			Endpoint: "https://synthetics.newrelic.com/synthetics/api/v3/monitors",
		}),
	}

	return svc
}

/*
Ping	SIMPLE
Simple browser	BROWSER
Scripted browser	SCRIPT_BROWSER
API test	SCRIPT_API
*/

/*
The Synthetics REST API limits an account's rate of requests to three requests per second. Requests made in excess of this threshold will return a 429 response code.
*/

// curl -v \
//  -H 'X-Api-Key:{Admin_User_Key}' https://synthetics.newrelic.com/synthetics/api/v3/monitors
/*
		Query arguments:

	offset: The monitor count offset. Defaults to 0. For example, if you have 40 monitors and you use an offset value of 20, it will return monitors 21-40.
	limit: The number of results per page, maximum 100. Defaults to 20.
*/

package syntehics

import (
	"github.com/RafPe/go-newrelic/core/client"
	request "github.com/RafPe/go-newrelic/core/nrq"
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
	// c := p.ClientConfig(EndpointsID, cfgs...)
	// if c.SigningNameDerived || len(c.SigningName) == 0 {
	// 	c.SigningName = "amplify"
	// }
	return newClient()
}

// newClient creates, initializes and returns a new service client instance.
func newClient() *Syntethics {
	svc := &Syntethics{}
	// svc := &Syntethics{
	// 	Client: client.New(
	// 		cfg,
	// 		metadata.ClientInfo{
	// 			ServiceName:   ServiceName,
	// 			ServiceID:     ServiceID,
	// 			SigningName:   signingName,
	// 			SigningRegion: signingRegion,
	// 			Endpoint:      endpoint,
	// 			APIVersion:    "2017-07-25",
	// 		},
	// 		handlers,
	// 	),
	// }

	// // Handlers
	// svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	// svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	// svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	// svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	// svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// // Run custom client initialization if present
	// if initClient != nil {
	// 	initClient(svc.Client)
	// }

	return svc
}

// newRequest creates a new request for a Amplify operation and runs any
// custom request initialization.
func (c *Syntethics) newRequest() *request.Nrq {
	req := &request.Nrq{}
	// req := c.NewRequest(op, params, data)

	// // Run custom request initialization if present
	// if initRequest != nil {
	// 	initRequest(req)
	// }

	return req
}

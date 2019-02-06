# New Relic API library for Go

This is a Go library that wraps the [New Relic][1] REST
API. It provides the needed types to interact with the New Relic REST API.

It's still in progress and I have finished only some parts of it.

The API documentation can be found from [New Relic][1],
and you'll need an API key (for some operations, an Admin API key is
required).

## USAGE

This library will provide a client object and any operations can be performed
through it. Simply import this library and create a client to get started:

```go
package main

```

## Contributing

Client is limited wihh functionality I really needed. Therefore if you like it and wanna contribute please fork and create PR.

**DISCLAIMER:** *I am in no way affiliated with New Relic and this work is
merely a convenience project for myself with no guarantees. It should be
considered "as-is" with no implication of responsibility. See the included
LICENSE for more details.*

[1]: http://www.newrelic.com
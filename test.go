package main

import (
	syntehics "github.com/RafPe/go-newrelic/services/syntethics"
)

func main() {

	x := syntehics.New()

	x.GetAllMonitors()

}

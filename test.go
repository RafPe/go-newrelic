package main

import (
	syntehics "github.com/RafPe/go-newrelic/services/syntethics"
)

func main() {

	x := syntehics.New()

	x.GetAllMonitors()

	x.GetMonitorByID("91d4a323-2dbd-xxx-995d-6663b44c5b62")

	x.CreateSimpleMonitor()

}

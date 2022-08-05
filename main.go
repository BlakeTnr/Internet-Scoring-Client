package main

import (
	"fmt"
	"isc/api"
)

func main() {
	var webServer api.WebServer
	address := api.Address{"", 23021}
	fmt.Println("Starting API on " + address.GetAddressString())
	webServer.Start(address.GetAddressString())
}

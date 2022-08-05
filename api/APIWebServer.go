package api

import (
	"fmt"
	routes "isc/api/routes"
	"net/http"
)

type WebServer struct {
	httpServer http.Server
	serveMux   http.ServeMux
}

// Multiple server don't work because then servers would have to have their own thread
func (webserver *WebServer) Start(address string) error {
	// THIS ENTIRE FUNCTION SHOULD BE REVIEWED AS IT IS MESSY AND CODE SHOULD BE DELEGATED TO OTHER OBJECTS
	webserver.RegisterPaths()
	webserver.httpServer = http.Server{Addr: address, Handler: &webserver.serveMux}

	err := webserver.httpServer.ListenAndServe()

	if err != nil {
		return fmt.Errorf("Couldn't start API webserver")
	}
	return nil
}

func (webserver *WebServer) Stop() {
	webserver.httpServer.Close()
}

func (webserver *WebServer) RegisterPaths() {
	webserver.serveMux.HandleFunc("/internetconnection", routes.VerifyInternetRouteHandler)
}

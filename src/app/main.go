package main

import (
	"configuration"
	"github.com/urfave/negroni"
	"google.golang.org/appengine"
	"handlers"
	"net/http"
)

func main() {
	routing := configuration.RoutingConfiguration{}

	ng := negroni.Classic()
	ng.Use(negroni.HandlerFunc(handlers.VerifyUser))
	ng.UseHandler(routing.GetRouter())
	http.Handle("/", ng)

	appengine.Main()
}

package main

import (
	"PlayingWithGoAndVue/configuration"
	"PlayingWithGoAndVue/handlers"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	routing := configuration.RoutingConfiguration{}

	ng := negroni.Classic()
	ng.Use(negroni.HandlerFunc(handlers.VerifyUser))
	ng.UseHandler(routing.GetRouter())
	http.Handle("/", ng)
	ng.Run(":8080")
}

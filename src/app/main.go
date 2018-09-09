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

	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//http.HandleFunc("/", handlers.RootHandler)
	//http.HandleFunc("/Earning", handlers.GetEarning)
	//http.HandleFunc("/Dividend", handlers.GetDividend)
	//http.HandleFunc("/dd44c44a-590f-4fba-bae0-aced5131fcd7/Import", handlers.Import)
	//http.HandleFunc("/login", handlers.LoginHandler)
	//http.HandleFunc("/login-ui", handlers.LoginUIHandler)
	//http.HandleFunc("/register-ui", handlers.RegisterUIHandler)
	//http.HandleFunc("/register", handlers.RegisterHandler)
	//http.HandleFunc("/check-username", handlers.UsernameAvailability)
	//http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
	//	http.ServeFile(writer, request, "./static/favicon.ico")
	//})
	appengine.Main()
}

package main

import (
	_ "./src/github.com/mattn/go-sqlite3"
	"./src/github.com/urfave/negroni"
	"app/configuration"
	"app/handlers"
	"log"
	"net/http"
)

func main() {
	ng := negroni.Classic()
	ng.Use(negroni.HandlerFunc(handlers.VerifyUser))
	ng.UseHandler(configuration.GetRouter())
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

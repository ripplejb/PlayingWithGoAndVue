package main

import (
	_ "./src/github.com/mattn/go-sqlite3"
	"./src/github.com/urfave/negroni"
	"app/configuration"
	"log"
	"net/http"
)

func main() {
	ng := negroni.Classic()
	ng.UseHandler(configuration.GetRouter())
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

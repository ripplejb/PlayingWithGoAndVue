package main

import (
	_ "./src/github.com/mattn/go-sqlite3"
	"./src/github.com/urfave/negroni"
	. "app/handlers"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/getEarnings", GetEarningsHandler)
	mux.HandleFunc("/getDividends", GetDividendsHandler)
	mux.HandleFunc("/dd44c44a-590f-4fba-bae0-aced5131fcd7/downloadSymbols", DownloadSymbols)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates"))))

	ng := negroni.Classic()
	ng.UseHandler(mux)
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

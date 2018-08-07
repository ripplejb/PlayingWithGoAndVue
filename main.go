package main

import (
	"./src/github.com/codegangsta/negroni"
	_ "./src/github.com/mattn/go-sqlite3"
	. "app/handlers"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/searchbooks", SearchBooksHandler)
	mux.HandleFunc("/getEarnings", GetEarningsHandler)
	mux.HandleFunc("/books/add", AddBookHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates"))))

	ng := negroni.Classic()
	ng.UseHandler(mux)
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

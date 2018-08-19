package main

import (
	"./src/github.com/gorilla/mux"
	_ "./src/github.com/mattn/go-sqlite3"
	"./src/github.com/urfave/negroni"
	"app/handlers"
	"log"
	"net/http"
)

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.RootHandler).Methods("GET")
	mux.HandleFunc("/Earning", handlers.GetEarning).Methods("GET")
	mux.HandleFunc("/Dividend", handlers.GetDividend).Methods("GET")
	mux.HandleFunc("/dd44c44a-590f-4fba-bae0-aced5131fcd7/Import", handlers.Import).Methods("GET")
	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/"))))

	ng := negroni.Classic()
	ng.UseHandler(mux)
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

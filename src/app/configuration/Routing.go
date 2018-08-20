package configuration

import (
	"app/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.RootHandler).Methods("GET")
	mux.HandleFunc("/Earning", handlers.GetEarning).Methods("GET")
	mux.HandleFunc("/Dividend", handlers.GetDividend).Methods("GET")
	mux.HandleFunc("/dd44c44a-590f-4fba-bae0-aced5131fcd7/Import", handlers.Import).Methods("GET")
	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/"))))
	mux.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./templates/favicon.ico")
	})

	return mux
}

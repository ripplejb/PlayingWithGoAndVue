package configuration

import (
	"github.com/gorilla/mux"
	"handlers"
	"net/http"
)

type RoutingConfiguration struct {
}

func (rc *RoutingConfiguration) GetRouter() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.RootHandler).Methods("GET")
	mux.HandleFunc("/Earning", handlers.GetEarning).Methods("GET")
	mux.HandleFunc("/Dividend", handlers.GetDividend).Methods("GET")
	mux.HandleFunc("/dd44c44a-590f-4fba-bae0-aced5131fcd7/Import", handlers.Import).Methods("GET")
	mux.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	mux.HandleFunc("/login-ui", handlers.LoginUIHandler).Methods("GET")
	mux.HandleFunc("/register-ui", handlers.RegisterUIHandler).Methods("GET")
	mux.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	mux.HandleFunc("/check-username", handlers.UsernameAvailability).Methods("GET")
	mux.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./static/favicon.ico")
	})

	return mux
}

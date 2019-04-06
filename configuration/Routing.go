package configuration

import (
	"PlayingWithGoAndVue/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type RoutingConfiguration struct {
}

func (rc *RoutingConfiguration) GetRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("main/static"))))
	muxRouter.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./main//static/favicon.ico")
	})
	muxRouter.HandleFunc("/", handlers.RootHandler).Methods("GET")
	muxRouter.HandleFunc("/Earning", handlers.GetEarning).Methods("GET")
	muxRouter.HandleFunc("/Dividend", handlers.GetDividend).Methods("GET")
	muxRouter.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	muxRouter.HandleFunc("/login-ui", handlers.LoginUIHandler).Methods("GET")
	muxRouter.HandleFunc("/register-ui", handlers.RegisterUIHandler).Methods("GET")
	muxRouter.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	muxRouter.HandleFunc("/check-username", handlers.UsernameAvailability).Methods("GET")

	return muxRouter
}

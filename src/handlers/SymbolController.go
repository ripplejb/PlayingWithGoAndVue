package handlers

import (
	"PlayingWithGoAndVue/src/services"
	"encoding/json"
	"google.golang.org/appengine"
	"net/http"
)

func Import(w http.ResponseWriter, r *http.Request) {
	var result bool
	var err error

	symbolService := services.SymbolService{}
	if result, err = symbolService.Import(appengine.NewContext(r)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

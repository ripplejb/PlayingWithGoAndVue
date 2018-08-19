package handlers

import (
	. "app/services"
	"encoding/json"
	"net/http"
)

func Import(w http.ResponseWriter, r *http.Request) {
	var result bool
	var err error

	if result, err = LoadSymbols(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

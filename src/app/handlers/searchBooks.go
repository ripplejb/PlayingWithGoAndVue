package handlers

import (
	"app/models"
	. "app/services"
	"encoding/json"
	"net/http"
)

func SearchBooksHandler(w http.ResponseWriter, r *http.Request) {
	var result models.SearchResultView
	var err error

	if result, err = SearchClassifyAPI(r.URL.Query().Get("searchInput")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

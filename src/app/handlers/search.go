package handlers

import (
	"app/models"
	. "app/services"
	"encoding/json"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var results []models.SearchResult
	var err error

	if results, err = Search(r.URL.Query().Get("searchInput")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := models.SearchResultView{Results: results, Headers: []string{"Title", "Author", "Year", "ID"}, ColumnWidths: []string{"40%", "30%", "10%", "20%"}}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

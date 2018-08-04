package handlers

import (
	. "app/models"
	. "app/repository"
	. "app/services"
	"encoding/json"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var result SearchResult
	err := decoder.Decode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		var book ClassifyBookResponse
		if book, err = Find(result.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = InsertBook(book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

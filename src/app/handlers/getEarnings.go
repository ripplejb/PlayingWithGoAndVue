package handlers

import (
	"app/models"
	. "app/services"
	"encoding/json"
	"net/http"
)

func GetEarningsHandler(w http.ResponseWriter, r *http.Request) {
	var result models.EarningsView
	var err error

	if result, err = GetEarningsData(r.URL.Query().Get("symbol")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

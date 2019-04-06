package handlers

import (
	"PlayingWithGoAndVue/models"
	"PlayingWithGoAndVue/services"
	"encoding/json"
	"net/http"
)

func GetEarning(w http.ResponseWriter, r *http.Request) {
	var result models.EarningsView
	var err error

	earningService := services.EarningService{}
	if result, err = earningService.GetEarningsData(r.URL.Query().Get("symbol")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

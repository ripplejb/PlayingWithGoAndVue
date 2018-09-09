package handlers

import (
	"PlayingWithGoAndVue/src/services"
	"encoding/json"
	"google.golang.org/appengine"
	"models"
	"net/http"
)

func GetEarning(w http.ResponseWriter, r *http.Request) {
	var result models.EarningsView
	var err error

	earningService := services.EarningService{}
	if result, err = earningService.GetEarningsData(appengine.NewContext(r), r.URL.Query().Get("symbol")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

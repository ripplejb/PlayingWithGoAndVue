package handlers

import (
	"PlayingWithGoAndVue/models"
	"PlayingWithGoAndVue/services"
	"encoding/json"
	"google.golang.org/appengine"
	"net/http"
)

func GetDividend(w http.ResponseWriter, r *http.Request) {
	var result models.DividendsView
	var err error

	dividendService := services.DividendService{}
	if result, err = dividendService.GetDividendsData(appengine.NewContext(r), r.URL.Query().Get("symbol")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package services

import (
	. "app/clients"
	"app/models"
	"app/repository"
	"encoding/json"
)

func LoadSymbols() (bool, error) {
	var body []byte
	var err error

	if body, err = RestClientGet("https://api.iextrading.com/1.0/ref-data/symbols"); err != nil {
		return false, err
	}

	var e []models.StockSymbol
	err = json.Unmarshal(body, &e)

	var repo = repository.StockSymbolRepository{}
	if len(e) > 0 {
		err = repo.Insert(e)
	}

	return err == nil, err
}

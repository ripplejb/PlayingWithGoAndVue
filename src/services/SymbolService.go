package services

import (
	"clients"
	"context"
	"encoding/json"
	"models"
	"repository"
)

type SymbolService struct {
}

func (ss *SymbolService) Import(context context.Context) (bool, error) {
	var body []byte
	var err error

	client := clients.Client{}
	if body, err = client.RestClientGet(context, "https://api.iextrading.com/1.0/ref-data/symbols"); err != nil {
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

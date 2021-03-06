package services

import (
	"PlayingWithGoAndVue/clients"
	"PlayingWithGoAndVue/models"
	"encoding/json"
)

type DividendService struct {
}

func (ds *DividendService) GetDividendsData(symbol string) (models.DividendsView, error) {
	var body []byte
	var err error

	client := clients.Client{}
	if body, err = client.RestClientGet("https://api.iextrading.com/1.0/stock/" + symbol + "/dividends/5y"); err != nil {
		return models.DividendsView{}, err
	}

	var e []models.Dividends
	err = json.Unmarshal(body, &e)

	result := models.DividendsView{Dividends: e, Headers: []string{"Ex StockDate", "Payment StockDate", "Amount", "Indicated", "Flag"}, ColumnWidths: []string{"25%", "25%", "20%", "20%", "10%"}}

	return result, err

}

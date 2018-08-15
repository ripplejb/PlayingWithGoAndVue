package models

type StockSymbol struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	StockDate string `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	StockType string `json:"type"`
	IexID     string `json:"iexId"`
}

package models

type StockSymbol struct {
	Symbol    string `json:"symbol" db:"Symbol"`
	Name      string `json:"name" db:"Name"`
	StockDate string `json:"date" db:"StockDate"`
	IsEnabled bool   `json:"isEnabled" db:"IsEnabled"`
	StockType string `json:"type" db:"StockType"`
	IexID     string `json:"iexId" db:"IexID"`
	ID        int64  `db:"ID"`
}

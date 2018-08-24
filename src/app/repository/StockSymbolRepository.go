package repository

import (
	. "app/dbproviders"
	"app/models"
)

type StockSymbolRepository struct {
}

func (s *StockSymbolRepository) Insert(symbols []models.StockSymbol) error {
	dbmap, err := GetSqlite3Database()

	if err == nil {
		_, err = dbmap.Exec("delete from StockSymbol")
		if err == nil {
			for _, element := range symbols {
				element.ID = -1
				if err = dbmap.Insert(&element); err == nil {
					//TODO: Handle error here
				}
			}
		}
	}

	return err
}

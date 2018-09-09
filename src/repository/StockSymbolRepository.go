package repository

import (
	"PlayingWithGoAndVue/src/dbproviders"
	"models"
)

type StockSymbolRepository struct {
}

func (s *StockSymbolRepository) Insert(symbols []models.StockSymbol) error {
	provider := dbproviders.Sqlite3Provider{}
	dbmap, err := provider.GetSqlite3Database()

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

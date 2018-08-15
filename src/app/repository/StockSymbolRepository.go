package repository

import (
	. "app/dbproviders"
	"app/models"
)

func Insert(symbols []models.StockSymbol) error {
	db, err := GetSqlite3Database()

	if err == nil {
		_, err = db.Exec("delete from StockSymbols")
		if err == nil {
			for _, element := range symbols {
				_, err = db.Exec("insert into StockSymbols (Symbol, Name, Date, IsEnabled, Type, IexID) values "+
					"(?, ?, ?, ?, ?, ?)",
					element.Symbol, element.Name, element.StockDate,
					element.IsEnabled, element.StockType, element.IexID)
			}
		}
	}

	return err
}

package repository

import (
	. "app/dbproviders"
)

func Insert(symbol string) error {
	db, err := GetSqlite3Database()

	if err == nil {
		_, err = db.Exec("insert into StockSymbols (symbol) values "+
			"(?)", symbol)
	}

	return err
}

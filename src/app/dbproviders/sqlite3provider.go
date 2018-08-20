package dbproviders

import (
	"app/models"
	"database/sql"
	"github.com/go-gorp/gorp"
)

func GetSqlite3Database() (*gorp.DbMap, error) {

	var db *sql.DB

	db, _ = sql.Open("sqlite3", "dev.db")

	err := db.Ping()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(models.StockSymbol{}, "StockSymbol").SetKeys(true, "ID")
	dbmap.CreateTablesIfNotExists()

	return dbmap, err
}

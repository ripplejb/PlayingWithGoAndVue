package dbproviders

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"models"
)

type Sqlite3Provider struct {
}

func (sp *Sqlite3Provider) GetSqlite3Database() (*gorp.DbMap, error) {

	var db *sql.DB

	db, _ = sql.Open("sqlite3", "dev.db")

	err := db.Ping()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(models.StockSymbol{}, "StockSymbol").SetKeys(true, "ID")
	dbmap.AddTableWithName(models.User{}, "Users").SetKeys(false, "username")
	dbmap.CreateTablesIfNotExists()

	return dbmap, err
}

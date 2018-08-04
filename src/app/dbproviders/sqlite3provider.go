package dbproviders

import (
	"database/sql"
)

func GetSqlite3Database() (*sql.DB, error) {

	var db *sql.DB

	db, _ = sql.Open("sqlite3", "dev.db")

	err := db.Ping()

	return db, err
}

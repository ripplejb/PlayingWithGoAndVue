package repository

import (
	. "app/dbproviders"
	. "app/models"
)

func InsertBook(book ClassifyBookResponse) error {
	db, err := GetSqlite3Database()

	if err == nil {
		_, err = db.Exec("insert into books (pk, title, author, id, classification) values "+
			"(?, ?, ?, ?, ?)", nil, book.BookData.Title, book.BookData.Author, book.BookData.Id, book.Classification.MostPopular)
	}

	return err
}

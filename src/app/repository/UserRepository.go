package repository

import (
	. "app/dbproviders"
	"app/models"
)

type UserRepository struct {
}

func (u *UserRepository) Insert(user models.User) error {
	dbmap, err := GetSqlite3Database()

	if err == nil {
		err = dbmap.Insert(&user)
	}
	return err
}

func (u *UserRepository) Search(user models.User) (models.User, error) {
	dbmap, err := GetSqlite3Database()

	if err == nil {
		var result models.User
		err = dbmap.SelectOne(&result, "select * from Users where username=?", user.Username)
		return result, err
	}

	return models.User{}, err
}

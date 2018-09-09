package repository

import (
	"dbproviders"
	"models"
)

type UserRepository struct {
}

func (u *UserRepository) Insert(user models.User) error {
	provider := dbproviders.Sqlite3Provider{}
	dbmap, err := provider.GetSqlite3Database()

	if err == nil {
		err = dbmap.Insert(&user)
	}
	return err
}

func (u *UserRepository) Search(user models.User) (models.User, error) {
	provider := dbproviders.Sqlite3Provider{}
	dbmap, err := provider.GetSqlite3Database()

	if err == nil {
		if found, err := dbmap.Get(models.User{}, user.Username); err == nil {
			result, ok := found.(*models.User)
			if ok {
				return *result, err
			}
		}
	}
	return models.User{}, err
}

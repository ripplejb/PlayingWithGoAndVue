package repository

import (
	"PlayingWithGoAndVue/dataaccesslayer"
	"PlayingWithGoAndVue/models"
)

type UserRepository struct {
	UserDB dataaccesslayer.IUserDB
}

func (u *UserRepository) Insert(user *models.User) error {
	err := u.UserDB.Add(user)
	return err
}

func (u *UserRepository) Search(user *models.User) (models.User, error) {
	usr, err := u.UserDB.Get(user.Username)

	return usr, err
}

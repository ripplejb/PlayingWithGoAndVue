package services

import (
	"app/models"
	"app/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
}

func (ls *LoginService) Authenticate(userename string, password string) error {
	repo := &repository.UserRepository{}
	user := models.User{Username: userename}
	user.SetSecret(password)

	result, err := repo.Search(user)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(result.Secret), []byte(password))
	}
	return err
}

func (ls *LoginService) Register(user models.User) error {
	repo := &repository.UserRepository{}
	err := repo.Insert(user)
	return err
}

func (ls *LoginService) CheckUserAvailability(username string) (bool, error) {
	repo := &repository.UserRepository{}
	result, err := repo.Search(models.User{Username: username})

	if err != nil {
		return false, err
	}
	return result.Username == username, nil
}

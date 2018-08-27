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

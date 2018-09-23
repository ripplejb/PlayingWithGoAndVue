package dataaccesslayer

import "models"

type IUserDB interface {
	Add(user *models.User) error
	Get(username string) (models.User, error)
}

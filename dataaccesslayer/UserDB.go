package dataaccesslayer

import (
	"PlayingWithGoAndVue/models"
	"context"
	"google.golang.org/appengine/datastore"
)

type UserDB struct {
	Context context.Context
}

func (u *UserDB) Add(user *models.User) error {
	key := datastore.NewIncompleteKey(u.Context, "User", nil)
	if _, err := datastore.Put(u.Context, key, user); err != nil {
		return err
	}
	return nil
}

func (u *UserDB) Get(username string) (models.User, error) {
	q := datastore.NewQuery("User").Filter("Username =", username)
	t := q.Run(u.Context)
	var user models.User
	_, err := t.Next(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

var _ IUserDB = &UserDB{}

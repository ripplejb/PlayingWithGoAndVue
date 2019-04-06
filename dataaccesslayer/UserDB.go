package dataaccesslayer

import (
	"PlayingWithGoAndVue/models"
	"cloud.google.com/go/datastore"
	"context"
)

type UserDB struct {
}

func GetContextClient() (ctx context.Context, client *datastore.Client) {
	ctx = context.Background()
	client, err := datastore.NewClient(ctx, "mycloud-169422")
	if err != nil {
		panic(err)
	}
	return ctx, client
}

func (u *UserDB) Add(user *models.User) error {
	ctx, client := GetContextClient()
	key := datastore.IncompleteKey("User", nil)
	if _, err := client.Put(ctx, key, user); err != nil {
		return err
	}
	return nil
}

func (u *UserDB) Get(username string) (models.User, error) {
	ctx, client := GetContextClient()

	q := datastore.NewQuery("User").Filter("Username =", username)
	t := client.Run(ctx, q)
	var user models.User
	_, err := t.Next(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

var _ IUserDB = &UserDB{}

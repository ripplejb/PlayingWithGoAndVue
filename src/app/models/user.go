package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username" db:"username"`
	secret   []byte `db:"secret"`
}

func (u *User) SetSecret(value string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err == nil {
		u.secret = bytes
	}
}

func (u User) Secret() []byte {
	return u.secret
}

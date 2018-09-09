package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username" db:"username"`
	Secret   []byte `db:"secret"`
}

func (u *User) SetSecret(value string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err == nil {
		u.Secret = bytes
	}
}

func (u User) GetSecret() []byte {
	return u.Secret
}

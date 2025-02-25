package user

import (
	"errors"
	"time"
)

type Admin struct {
	email    string
	password string
	User
}

func (a *Admin) New(email, password, firstName, lastName, birthday string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("all fields are required")
	}
	return &Admin{email, password, User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}}, nil
}

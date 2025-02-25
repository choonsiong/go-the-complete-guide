package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
	updatedAt time.Time
}

// newUser returns a new *User value
func newUser(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// setFirstName set user first name
func (u *User) setFirstName(firstName string) {
	(*u).firstName = firstName
}

// setLastName set user last name
func (u *User) setLastName(lastName string) {
	(*u).lastName = lastName
}

// setBirthday set user birthday
func (u *User) setBirthday(birthday string) {
	u.birthday = birthday
}

// description returns a string describing user's details
func (u *User) description() string {
	return fmt.Sprintf("First name: %s\nLast name:%s\nBirthday: %s\nCreated: %s\nUpdated: %s", u.firstName, u.lastName, u.birthday, u.createdAt, u.updatedAt)
}

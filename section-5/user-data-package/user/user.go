package user

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

// New returns a new *User value
func New(firstName, lastName, birthday string) (*User, error) {
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

// SetFirstName set user first name
func (u *User) SetFirstName(firstName string) {
	(*u).firstName = firstName
}

// SetLastName set user last name
func (u *User) SetLastName(lastName string) {
	(*u).lastName = lastName
}

// SetBirthday set user birthday
func (u *User) SetBirthday(birthday string) {
	u.birthday = birthday
}

// Description returns a string consists of all User fields
func (u *User) Description() string {
	return fmt.Sprintf("First name: %s\nLast name:%s\nBirthday: %s\nCreated: %s\nUpdated: %s", u.firstName, u.lastName, u.birthday, u.createdAt, u.updatedAt)
}

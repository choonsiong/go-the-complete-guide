package models

import (
	"errors"
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Save insert a new user to the database or returns an error if any
func (u User) Save() (*User, error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	encryptedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	if encryptedPassword == "" {
		return nil, errors.New("encrypted password is empty")
	}

	result, err := stmt.Exec(u.Email, encryptedPassword)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = int(id)

	return &u, nil
}

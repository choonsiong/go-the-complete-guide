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

func (u *User) ValidateCredential() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var userID int
	var retrievedPassword string
	err := row.Scan(&userID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials 😜")
	}

	if !utils.IsPasswordValid(retrievedPassword, u.Password) {
		return errors.New("invalid credentials 😜")
	}

	u.ID = userID

	return nil
}

// Save insert a new user to the database or returns an error if any
func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	encryptedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	if encryptedPassword == "" {
		return errors.New("encrypted password is empty")
	}

	result, err := stmt.Exec(u.Email, encryptedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)

	return nil
}

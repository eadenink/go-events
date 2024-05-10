package models

import (
	"github.com/eadenink/go-events/db"
	"github.com/eadenink/go-events/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = userId

	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	err = utils.VerifyPassword(retrievedPassword, user.Password)
	return err
}

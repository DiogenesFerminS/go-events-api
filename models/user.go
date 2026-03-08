package models

import (
	"errors"

	"go_event_api.com/go_api/db"
	"go_event_api.com/go_api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (User, error) {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return User{}, err
	}

	u.Password = hashedPassword

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return User{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return User{}, err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return User{}, err
	}

	u.ID = userId
	return u, nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	isValidPassword := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isValidPassword {
		return errors.New("Invalid credentials")
	}

	return nil

}

package services

import (
	"errors"
	"event-booking-api/db"
	"event-booking-api/models"
)

func AddUser(user *models.User) (int64, error) {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	if err := user.EncodePassword(); err != nil {
		return 0, err
	}
	result, err := db.Connection.Exec(query, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func LoginUser(user *models.User) (string, error) {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.Connection.QueryRow(query, user.Email)
	var userId int64
	var password string
	if err := row.Scan(&userId, &password); err != nil {
		return "", errors.New("invalid credentials")
	}

	if !user.ComparePassword(password) {
		return "", errors.New("invalid credentials")
	}
	user.Id = userId
	return user.GenerateToken()
}

package models

import (
	conn "MarketPlaceBackEnd/internal/database"
	crypt "MarketPlaceBackEnd/tools"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func CreateUser(email string, pass string) error {
	db, err := conn.ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// Hashing user password
	hash, err := crypt.HashingPassword(pass)
	if err != nil {
		logrus.Warnf("Error hashing password - %s", err)
		return err
	}

	// Adding a new record to a table with users
	if _, err := db.Exec(`INSERT INTO users (email, hash_pass) VALUES($1, $2)`, email, hash); err != nil {
		logrus.Warnf("Error creating a new user - %s", err)
		return err
	}
	logrus.Info("A new user was successfully added")
	return nil
}

func AuthUser(email string, pass string) (bool, error) {
	var temp string
	db, err := conn.ConnectDatabase()
	if err != nil {
		return false, err
	}
	defer db.Close()

	if rows := db.QueryRow(`SELECT hash_pass FROM users WHERE email = $1`, email).Scan(&temp); rows == sql.ErrNoRows {
		return false, nil
	}

	if crypt.CheckControlSum(temp, pass) {
		return false, nil
	}

	logrus.Info("The user was successfully logged in")
	return true, nil
}

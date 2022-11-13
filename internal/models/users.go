package models

import (
	conn "MarketPlaceBackEnd/internal/database"
	crypt "MarketPlaceBackEnd/tools"

	"github.com/sirupsen/logrus"
)

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func CreateUser(email string, pass string) error {
	// Connect to database
	db, err := conn.ConnectDatabase()
	if err != nil {
		return err
	}

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

func AuthUser() {

}

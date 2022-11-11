package database

import (
	"database/sql"
	"fmt"

	conf "MarketPlaceBackEnd/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectDatabase() (*sql.DB, error) {
	var cfg conf.ConfigDb
	username, host, port, password, dbname, sslmode := cfg.ReadConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbname, sslmode))
	if err != nil {
		logrus.Warnf("Err open connect to database - %s", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.Warnf("Err check connect - %s", err)
		return nil, err
	}
	logrus.Info("Connection to the database was successful")

	return db, nil
}

package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ConfigDb struct {
	Database struct {
		Username string `yaml:"username"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func (db *ConfigDb) ReadConfig() (string, string, string, string, string, string) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Fatalf("Err read config.yaml file - %s", err)
	}

	if err := yaml.Unmarshal(file, db); err != nil {
		logrus.Fatalf("Err unmarshal config.yaml to struct - %s", err)
	}
	return db.Database.Username, db.Database.Host, db.Database.Port, db.Database.Password, db.Database.Dbname, db.Database.Sslmode
}

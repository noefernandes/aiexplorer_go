package configs

import (
	"os"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func Load() error {
	cfg = new(config)

	cfg.API = APIConfig{
		Port: os.Getenv("PORT"),
	}

	cfg.DB = DBConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Pass:     os.Getenv("DBPASS"),
		Database: os.Getenv("DBNAME"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

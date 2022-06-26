package config

import (
	"fmt"
	"go_db_fundamental/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Db *sqlx.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")         //"localhost"
	dbPort := os.Getenv("DB_PORT")         //"5432"
	dbUser := os.Getenv("DB_USER")         //"postgres"
	dbPassword := os.Getenv("DB_PASSWORD") //"12345678"
	dbName := os.Getenv("DB_NAME")         //"db_enigma_shop"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dsn)

	utils.IsError(err)

	c.Db = db

}

func (c *Config) DbConn() *sqlx.DB {
	return c.Db
}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}

package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Db *sql.DB
}

func (c *Config) initDb() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "P@ssw0rd"
	dbName := "db_sinar_harapan_makmur"
	dbDriver := "postgres"

	// dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, user, password, host, port, dbName)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open(dbDriver, psqlInfo)
	if err != nil {
		panic(err)
	}
	c.Db = db
}

func (c *Config) DbConn() *sql.DB {
	return c.Db
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}

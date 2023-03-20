package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DbConection interface {
	Conn() *sql.DB
}

type dbConnection struct {
	db  *sql.DB
	cfg Config
}

func (d *dbConnection) initDb() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.cfg.Host, d.cfg.Port, d.cfg.User, d.cfg.Password, d.cfg.Name)
	db, err := sql.Open(d.cfg.Driver, psqlconn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to run", err)
		}
	}()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	d.db = db
	log.Println("Database connected")
}

func (d *dbConnection) Conn() *sql.DB {
	return d.db
}

func NewDbConnection(config Config) DbConection {
	infra := dbConnection{
		cfg: config,
	}
	infra.initDb()
	return &infra
}

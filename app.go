package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
	_ "github.com/lib/pq"
)

func main() {

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "P@ssw0rd"
	dbname := "db_sinar_harapan_makmur"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Driver tidak ditemukan : %s", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	figure := figure.NewFigure("Enigma Camp", "standard", true)
	figure.Print()
	fmt.Println("Successfully connected!")
}

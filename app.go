package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
)

func main() {

	c := config.NewConfig()
	dbConn := config.NewDbConnection(c)

	defer dbConn.Conn().Close()

	err := dbConn.Conn().Ping()
	if err != nil {
		panic(err)
	}

	figure := figure.NewFigure("Enigma Camp", "standard", true)
	figure.Print()
	fmt.Println("Successfully connected!")
}

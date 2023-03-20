package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
)

func main() {

	cfg := config.NewConfig()
	defer cfg.DbConn().Close()

	err := cfg.DbConn().Ping()
	if err != nil {
		panic(err)
	}

	figure := figure.NewFigure("Enigma Camp", "standard", true)
	figure.Print()
	fmt.Println("Successfully connected!")
}

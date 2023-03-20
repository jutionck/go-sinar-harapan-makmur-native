package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
)

func main() {

	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)

	defer dbConn.Conn().Close()

	figure := figure.NewFigure("Enigma Camp", "standard", true)
	figure.Print()
	fmt.Println("Successfully connected!")
}

package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

func main() {
	vehicle := model.Vehicle{
		Brand:          "Honda",
		Model:          "Civic",
		ProductionYear: 2022,
		Color:          "Red",
		IsAutomatic:    true,
		Stock:          10,
		SalePrice:      250000000,
		Status:         "Baru",
	}
	if vehicle.IsValidStatus() {
		fmt.Println("Status is valid")
	} else {
		fmt.Println("Status is not valid")
		return
	}

	vehicle.SetId()
	fmt.Println(vehicle)

}

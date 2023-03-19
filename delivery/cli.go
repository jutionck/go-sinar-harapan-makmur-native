package delivery

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
)

func VehicleCLI() {
	vehicleRepository := repository.NewVehicleRepository()
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepository)

	// cretae
	newVehicle := model.Vehicle{
		Id:        "V0006",
		Brand:     "Toyota",
		Model:     "Alphard",
		Color:     "Putih",
		Stock:     0,
		Status:    "Baru",
		SalePrice: 1500000000000,
	}

	if err := vehicleUseCase.RegisterNewVehicle(newVehicle); err != nil {
		fmt.Println(err)
		return
	}

	// Get All
	vehicles, err := vehicleUseCase.FindAllVehicle()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range vehicles {
			fmt.Println("ID:", v.Id)
			fmt.Println("Brand:", v.Brand)
			fmt.Println("Model:", v.Model)
			fmt.Println("Production Year:", v.ProductionYear)
			fmt.Println("Color:", v.Color)
			fmt.Println("Price:", v.SalePrice)
			fmt.Println("Is Automatic:", v.IsAutomatic)
			fmt.Println("Stock:", v.Stock)
			fmt.Println("Status:", v.Status)
			fmt.Println()
		}
	}

}

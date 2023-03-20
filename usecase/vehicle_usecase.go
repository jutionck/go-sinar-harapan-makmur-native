package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type VehicleUseCase interface {
	RegisterNewVehicle(newVehicle model.Vehicle) error
	FindAllVehicle() ([]model.Vehicle, error)
	GetVehicle(id string) (model.Vehicle, error)
	UpdateVehicle(newVehicle model.Vehicle) error
	DeleteVehicle(id string) error
}

type vehicleUseCase struct {
	vehicleRepo repository.VehicleRepository
}

func (v *vehicleUseCase) RegisterNewVehicle(newVehicle model.Vehicle) error {
	isExists, _ := v.GetVehicle(newVehicle.Id)
	if isExists.Id == newVehicle.Id {
		return fmt.Errorf("Vehicle with ID: %v exists", newVehicle.Id)
	}

	if newVehicle.Brand == "" || newVehicle.Model == "" || newVehicle.Color == "" {
		return fmt.Errorf("Brand, Model, and Color are required fields")
	}

	if !newVehicle.IsValidStatus() {
		return fmt.Errorf("Invalid status: %s", newVehicle.Status)
	}

	if newVehicle.SalePrice < 0 || newVehicle.SalePrice == 0 {
		return fmt.Errorf("Sale price can't zero or negative ")
	}

	if newVehicle.Stock < 0 {
		return fmt.Errorf("Stock can't negative ")
	}

	err := v.vehicleRepo.Create(newVehicle)
	if err != nil {
		return fmt.Errorf("Failed to create new vehicle: %v", err)
	}

	return nil
}

func (v *vehicleUseCase) FindAllVehicle() ([]model.Vehicle, error) {
	return v.vehicleRepo.List()
}

func (v *vehicleUseCase) GetVehicle(id string) (model.Vehicle, error) {
	return v.vehicleRepo.Get(id)
}

func (v *vehicleUseCase) UpdateVehicle(newVehicle model.Vehicle) error {
	if newVehicle.Brand == "" || newVehicle.Model == "" || newVehicle.Color == "" {
		return fmt.Errorf("Brand, Model, and Color are required fields")
	}

	if !newVehicle.IsValidStatus() {
		return fmt.Errorf("Invalid status: %s", newVehicle.Status)
	}

	if newVehicle.SalePrice < 0 || newVehicle.SalePrice == 0 {
		return fmt.Errorf("Sale price can't zero or negative ")
	}

	if newVehicle.Stock < 0 {
		return fmt.Errorf("Stock can't negative ")
	}

	err := v.vehicleRepo.Update(newVehicle)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (v *vehicleUseCase) DeleteVehicle(id string) error {
	return v.vehicleRepo.Delete(id)
}

func NewVehicleUseCase(vehicleRepo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{vehicleRepo: vehicleRepo}
}

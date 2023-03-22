package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type VehicleUseCase interface {
	RegisterNewVehicle(newVehicle model.Vehicle) error
	FindAllVehicle() ([]model.Vehicle, error)
	GetVehicle(id string) (model.Vehicle, error)
	UpdateVehicle(newVehicle model.Vehicle) error
	DeleteVehicle(id string) error
	Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging)
}

type vehicleUseCase struct {
	vehicleRepo repository.VehicleRepository
}

func (v *vehicleUseCase) RegisterNewVehicle(newVehicle model.Vehicle) error {

	if err := vehicleValidation(newVehicle); err != nil {
		return err
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
	if err := vehicleValidation(newVehicle); err != nil {
		return err
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

func (v *vehicleUseCase) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging) {
	return v.vehicleRepo.Paging(requestQueryParams)
}

func vehicleValidation(payload model.Vehicle) error {
	if payload.Brand == "" || payload.Model == "" || payload.Color == "" {
		return fmt.Errorf("Brand, Model, and Color are required fields")
	}

	if !payload.IsValidStatus() {
		return fmt.Errorf("Invalid status: %s", payload.Status)
	}

	if payload.SalePrice < 0 || payload.SalePrice == 0 {
		return fmt.Errorf("Sale price can't zero or negative ")
	}

	if payload.Stock < 0 {
		return fmt.Errorf("Stock can't negative ")
	}
	return nil
}

func NewVehicleUseCase(vehicleRepo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{vehicleRepo: vehicleRepo}
}

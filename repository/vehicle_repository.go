package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type VehicleRepository interface {
	BaseRepository[model.Vehicle]
}

type vehicleRepository struct {
	db []model.Vehicle
}

func (v *vehicleRepository) Create(newData model.Vehicle) error {
	v.db = append(v.db, newData)
	if len(v.db) == 0 {
		return fmt.Errorf("Gagal menyimpan data")
	}
	return nil
}

func (v *vehicleRepository) List() ([]model.Vehicle, error) {
	if len(v.db) == 0 {
		return nil, fmt.Errorf("Database kosong")
	}
	return v.db, nil
}

func (v *vehicleRepository) Get(id string) (model.Vehicle, error) {
	for _, vehicle := range v.db {
		if vehicle.Id == id {
			return vehicle, nil
		}
	}
	return model.Vehicle{}, fmt.Errorf("Data tidak ditemukan")
}

func (v *vehicleRepository) Update(newData model.Vehicle) error {
	for i, vehicle := range v.db {
		if vehicle.Id == newData.Id {
			v.db[i] = newData
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func (v *vehicleRepository) Delete(id string) error {
	for i, vehicle := range v.db {
		if vehicle.Id == id {
			v.db = append(v.db[:i], v.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func NewVehicleRepository() VehicleRepository {
	vehicles := []model.Vehicle{
		{
			Id:             "V0001",
			Brand:          "Honda",
			Model:          "HR-V",
			ProductionYear: 2022,
			Color:          "Putih",
			SalePrice:      301000000,
			IsAutomatic:    true,
			Stock:          10,
			Status:         "Baru",
		},
		{
			Id:             "V0002",
			Brand:          "Honda",
			Model:          "Civic",
			ProductionYear: 2021,
			Color:          "Putih",
			SalePrice:      301000000,
			IsAutomatic:    true,
			Stock:          7,
			Status:         "Bekas",
		},
		{
			Id:             "V0003",
			Brand:          "Mitshubishi",
			Model:          "XPander",
			ProductionYear: 2022,
			Color:          "Hitam",
			SalePrice:      231750000,
			IsAutomatic:    false,
			Stock:          5,
			Status:         "Baru",
		},
		{
			Id:             "V0004",
			Brand:          "Toyota",
			Model:          "Rush",
			ProductionYear: 2021,
			Color:          "Hitam",
			SalePrice:      232000000,
			IsAutomatic:    true,
			Stock:          5,
			Status:         "Baru",
		},
		{
			Id:             "V0005",
			Brand:          "Mazda",
			Model:          "CX-3",
			ProductionYear: 2022,
			Color:          "Putih",
			SalePrice:      302000000,
			IsAutomatic:    true,
			Stock:          15,
			Status:         "Baru",
		},
	}

	return &vehicleRepository{db: vehicles}
}

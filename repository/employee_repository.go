package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
}

type employeeRepository struct {
	db []model.Employee
}

func (e *employeeRepository) Create(newData model.Employee) error {
	e.db = append(e.db, newData)
	if len(e.db) == 0 {
		return fmt.Errorf("Gagal menyimpan data")
	}
	return nil
}

func (e *employeeRepository) List() ([]model.Employee, error) {
	if len(e.db) == 0 {
		return nil, fmt.Errorf("Database kosong")
	}
	return e.db, nil
}

func (e *employeeRepository) Get(id string) (model.Employee, error) {
	for _, employee := range e.db {
		if employee.Id == id {
			return employee, nil
		}
	}
	return model.Employee{}, fmt.Errorf("Data tidak ditemukan")
}

func (e *employeeRepository) Update(newData model.Employee) error {
	for i, employee := range e.db {
		if employee.Id == newData.Id {
			e.db[i] = newData
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func (e *employeeRepository) Delete(id string) error {
	for i, employee := range e.db {
		if employee.Id == id {
			e.db = append(e.db[:i], e.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func NewEmployeeRepository() EmployeeRepository {
	employees := []model.Employee{
		{
			Id:        "EP001",
			FirstName: "Edo",
			LastName:  "Sensei",
		},
		{
			Id:        "EP002",
			FirstName: "Angga",
			LastName:  "Raditya",
			Manager:   &model.Employee{Id: "EP001"},
		},
		{
			Id:        "EP003",
			FirstName: "Joe",
			LastName:  "Andrey",
			Manager:   &model.Employee{Id: "EP001"},
		},
	}
	return &employeeRepository{db: employees}
}

package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	GetByEmail(email string) (model.Employee, error)
	GetByPhoneNumber(phoneNumber string) (model.Employee, error)
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

func (c *employeeRepository) GetByEmail(email string) (model.Employee, error) {
	for _, customer := range c.db {
		if customer.Email == email {
			return customer, nil
		}
	}
	return model.Employee{}, fmt.Errorf("Employee with email: %s don't exists", email)
}

func (c *employeeRepository) GetByPhoneNumber(phoneNumber string) (model.Employee, error) {
	for _, customer := range c.db {
		if customer.PhoneNumber == phoneNumber {
			return customer, nil
		}
	}
	return model.Employee{}, fmt.Errorf("Employee with phone number: %s don't exists", phoneNumber)
}

func NewEmployeeRepository() EmployeeRepository {
	employees := []model.Employee{
		{
			Id:          "EP001",
			FirstName:   "Edo",
			LastName:    "Sensei",
			Email:       "edo.sensei@gmail.com",
			PhoneNumber: "0878282829",
			Posisition:  "Manager",
			Salary:      25000000,
		},
		{
			Id:          "EP002",
			FirstName:   "Angga",
			LastName:    "Raditya",
			Email:       "angga.raditya@gmail.com",
			PhoneNumber: "0857282829",
			Posisition:  "Staff A",
			Salary:      85000000,
			Manager: &model.Employee{
				Id:          "EP001",
				FirstName:   "Edo",
				LastName:    "Sensei",
				Email:       "edo.sensei@gmail.com",
				PhoneNumber: "0878282829",
				Posisition:  "Manager",
				Salary:      25000000,
			},
		},
		{
			Id:          "EP003",
			FirstName:   "Joe",
			LastName:    "Andrey",
			Email:       "joe.andrey@gmail.com",
			PhoneNumber: "0812282829",
			Posisition:  "Staff B",
			Salary:      7500000,
			Manager: &model.Employee{
				Id:          "EP001",
				FirstName:   "Edo",
				LastName:    "Sensei",
				Email:       "edo.sensei@gmail.com",
				PhoneNumber: "0878282829",
				Posisition:  "Manager",
				Salary:      25000000,
			},
		},
	}
	return &employeeRepository{db: employees}
}

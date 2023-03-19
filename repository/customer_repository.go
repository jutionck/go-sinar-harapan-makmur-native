package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
}

type customerRepository struct {
	db []model.Customer
}

func (c *customerRepository) Create(newData model.Customer) error {
	c.db = append(c.db, newData)
	if len(c.db) == 0 {
		return fmt.Errorf("Gagal menyimpan data")
	}
	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	if len(c.db) == 0 {
		return nil, fmt.Errorf("Database kosong")
	}
	return c.db, nil
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	for _, customer := range c.db {
		if customer.Id == id {
			return customer, nil
		}
	}
	return model.Customer{}, fmt.Errorf("Data tidak ditemukan")
}

func (c *customerRepository) Update(newData model.Customer) error {
	for i, customer := range c.db {
		if customer.Id == newData.Id {
			c.db[i] = newData
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func (c *customerRepository) Delete(id string) error {
	for i, customer := range c.db {
		if customer.Id == id {
			c.db = append(c.db[:i], c.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Data tidak ditemukan")
}

func NewCustomerRepository() CustomerRepository {
	customers := []model.Customer{
		{
			Id:        "C0001",
			FirstName: "Jution",
			LastName:  "Candra",
		},
		{
			Id:        "C0002",
			FirstName: "Fadli",
			LastName:  "Rahman",
		},
		{
			Id:        "C0003",
			FirstName: "Tika",
			LastName:  "Yesi",
		},
	}
	return &customerRepository{db: customers}
}

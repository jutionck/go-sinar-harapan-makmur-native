package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	GetByEmail(email string) (model.Customer, error)
	GetByPhoneNumber(phoneNumber string) (model.Customer, error)
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

func (c *customerRepository) GetByEmail(email string) (model.Customer, error) {
	for _, customer := range c.db {
		if customer.Email == email {
			return customer, nil
		}
	}
	return model.Customer{}, fmt.Errorf("Customer with email: %s don't exists", email)
}

func (c *customerRepository) GetByPhoneNumber(phoneNumber string) (model.Customer, error) {
	for _, customer := range c.db {
		if customer.PhoneNumber == phoneNumber {
			return customer, nil
		}
	}
	return model.Customer{}, fmt.Errorf("Customer with phone number: %s don't exists", phoneNumber)
}

func NewCustomerRepository() CustomerRepository {
	customers := []model.Customer{
		{
			Id:          "C0001",
			FirstName:   "Jution",
			LastName:    "Candra",
			Email:       "jution.candra@gmail.com",
			PhoneNumber: "0821111111",
		},
		{
			Id:          "C0002",
			FirstName:   "Fadli",
			LastName:    "Rahman",
			Email:       "fadlo.rahman@gmail.com",
			PhoneNumber: "0821222222",
		},
		{
			Id:          "C0003",
			FirstName:   "Tika",
			LastName:    "Yesi",
			Email:       "tika.yesi@gmail.com",
			PhoneNumber: "0821333333",
		},
	}
	return &customerRepository{db: customers}
}

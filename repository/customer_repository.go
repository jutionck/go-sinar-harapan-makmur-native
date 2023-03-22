package repository

import (
	"database/sql"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	GetByEmail(email string) (model.Customer, error)
	GetByPhoneNumber(phoneNumber string) (model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(newData model.Customer) error {
	sql := "INSERT INTO customer (id, first_name, last_name, address, phone_number, email, bod) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := c.db.Exec(sql, newData.Id, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod FROM customer`
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address, &customer.PhoneNumber, &customer.Email, &customer.Bod)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod FROM customer WHERE id = $1`
	var customer model.Customer
	err := c.db.QueryRow(sql, id).Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address, &customer.PhoneNumber, &customer.Email, &customer.Bod)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) Update(newData model.Customer) error {
	sql := "UPDATE customer SET first_name = $1, last_name = $2, address = $3, phone_number = $4, email = $5, bod = $6 WHERE id = $7"
	_, err := c.db.Exec(sql, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod, newData.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id string) error {
	sql := "DELETE FROM customer WHERE id = $1"
	_, err := c.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) GetByEmail(email string) (model.Customer, error) {
	sql := `SELECT id, email FROM customer WHERE email = $1`
	var customer model.Customer
	err := c.db.QueryRow(sql, email).Scan(&customer.Id, &customer.Email)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) GetByPhoneNumber(phoneNumber string) (model.Customer, error) {
	sql := `SELECT id, phone_number FROM customer WHERE phone_number = $1`
	var customer model.Customer
	err := c.db.QueryRow(sql, phoneNumber).Scan(&customer.Id, &customer.PhoneNumber)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

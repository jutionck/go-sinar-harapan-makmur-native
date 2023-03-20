package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(newCustomer model.Customer) error
	FindAllCustomer() ([]model.Customer, error)
	GetCustomer(id string) (model.Customer, error)
	UpdateCustomer(newCustomer model.Customer) error
	DeleteCustomer(id string) error
	FindCustomerByEmail(email string) (model.Customer, error)
	FindCustomerByPhoneNumber(phoneNumber string) (model.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(newCustomer model.Customer) error {
	isExists, _ := c.GetCustomer(newCustomer.Id)
	if isExists.Id == newCustomer.Id {
		return fmt.Errorf("Customer with ID: %v exists", newCustomer.Id)
	}

	isEmailExist, _ := c.customerRepo.GetByEmail(newCustomer.Email)
	if isEmailExist.Email == newCustomer.Email {
		return fmt.Errorf("Customer with email: %v exists", newCustomer.Email)
	}

	isPhoneNumberExist, _ := c.customerRepo.GetByPhoneNumber(newCustomer.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newCustomer.PhoneNumber {
		return fmt.Errorf("Customer with phone number: %v exists", newCustomer.PhoneNumber)
	}

	if newCustomer.FirstName == "" || newCustomer.LastName == "" || newCustomer.PhoneNumber == "" || newCustomer.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	err := c.customerRepo.Create(newCustomer)
	if err != nil {
		return fmt.Errorf("Failed to create new vehicle: %v", err)
	}

	return nil
}

func (c *customerUseCase) FindAllCustomer() ([]model.Customer, error) {
	return c.customerRepo.List()
}

func (c *customerUseCase) GetCustomer(id string) (model.Customer, error) {
	return c.customerRepo.Get(id)
}

func (c *customerUseCase) UpdateCustomer(newCustomer model.Customer) error {
	isEmailExist, _ := c.customerRepo.GetByEmail(newCustomer.Email)

	if isEmailExist.Email == newCustomer.Email && isEmailExist.Id != newCustomer.Id {
		return fmt.Errorf("Customer with email: %v exists", newCustomer.Email)
	}

	isPhoneNumberExist, _ := c.customerRepo.GetByPhoneNumber(newCustomer.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newCustomer.PhoneNumber && isPhoneNumberExist.Id != newCustomer.Id {
		return fmt.Errorf("Customer with phone number: %v exists", newCustomer.PhoneNumber)
	}

	if newCustomer.FirstName == "" || newCustomer.LastName == "" || newCustomer.PhoneNumber == "" || newCustomer.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	err := c.customerRepo.Update(newCustomer)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.customerRepo.Delete(id)
}

func (c *customerUseCase) FindCustomerByEmail(email string) (model.Customer, error) {
	return c.customerRepo.GetByEmail(email)
}
func (c *customerUseCase) FindCustomerByPhoneNumber(phoneNumber string) (model.Customer, error) {
	return c.customerRepo.GetByPhoneNumber(phoneNumber)
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: customerRepo}
}

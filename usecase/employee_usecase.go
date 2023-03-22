package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(newEmployee model.Employee) error
	FindAllEmployee() ([]model.Employee, error)
	GetEmployee(id string) (model.Employee, error)
	UpdateEmployee(newEmployee model.Employee) error
	DeleteEmployee(id string) error
	FindEmployeeByEmail(email string) (model.Employee, error)
	FindEmployeeByPhoneNumber(phoneNumber string) (model.Employee, error)
	FindManagerById(id string) (model.Employee, error)
}

type employeeUseCase struct {
	employeeRepo repository.EmployeeRepository
}

func (e *employeeUseCase) RegisterNewEmployee(newEmployee model.Employee) error {
	isExists, _ := e.GetEmployee(newEmployee.Id)
	if isExists.Id == newEmployee.Id {
		return fmt.Errorf("Employee with ID: %v exists", newEmployee.Id)
	}

	isEmailExist, _ := e.employeeRepo.GetByEmail(newEmployee.Email)
	if isEmailExist.Email == newEmployee.Email {
		return fmt.Errorf("Employee with email: %v exists", newEmployee.Email)
	}

	isPhoneNumberExist, _ := e.employeeRepo.GetByPhoneNumber(newEmployee.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newEmployee.PhoneNumber {
		return fmt.Errorf("Employee with phone number: %v exists", newEmployee.PhoneNumber)
	}

	if newEmployee.FirstName == "" || newEmployee.LastName == "" || newEmployee.PhoneNumber == "" || newEmployee.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	if newEmployee.Manager != nil {
		manager, err := e.FindManagerById(newEmployee.Manager.Id)
		if err != nil {
			return fmt.Errorf("Manager with ID: %v exists", newEmployee.Manager.Id)
		}
		newEmployee.Manager = &manager
	} else {
		newEmployee.Manager = nil
	}

	err := e.employeeRepo.Create(newEmployee)
	if err != nil {
		return fmt.Errorf("Failed to create new employee: %v", err)
	}

	return nil
}

func (e *employeeUseCase) FindAllEmployee() ([]model.Employee, error) {
	return e.employeeRepo.List()
}

func (e *employeeUseCase) GetEmployee(id string) (model.Employee, error) {
	return e.employeeRepo.Get(id)
}

func (e *employeeUseCase) UpdateEmployee(newEmployee model.Employee) error {
	isEmailExist, _ := e.employeeRepo.GetByEmail(newEmployee.Email)

	if isEmailExist.Email == newEmployee.Email && isEmailExist.Id != newEmployee.Id {
		return fmt.Errorf("Employee with email: %v exists", newEmployee.Email)
	}

	isPhoneNumberExist, _ := e.employeeRepo.GetByPhoneNumber(newEmployee.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newEmployee.PhoneNumber && isPhoneNumberExist.Id != newEmployee.Id {
		return fmt.Errorf("Employee with phone number: %v exists", newEmployee.PhoneNumber)
	}

	if newEmployee.FirstName == "" || newEmployee.LastName == "" || newEmployee.PhoneNumber == "" || newEmployee.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	manager, err := e.FindManagerById(newEmployee.Manager.Id)
	if err != nil {
		return fmt.Errorf("Manager with ID: %v exists", newEmployee.Manager.Id)
	}

	if newEmployee.Manager.Id != "" {
		newEmployee.Manager = &manager
	}

	err = e.employeeRepo.Update(newEmployee)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	return e.employeeRepo.Delete(id)
}

func (e *employeeUseCase) FindEmployeeByEmail(email string) (model.Employee, error) {
	return e.employeeRepo.GetByEmail(email)
}

func (e *employeeUseCase) FindEmployeeByPhoneNumber(phoneNumber string) (model.Employee, error) {
	return e.employeeRepo.GetByPhoneNumber(phoneNumber)
}

func (e *employeeUseCase) FindManagerById(id string) (model.Employee, error) {
	return e.GetEmployee(id)
}

func NewEmployeeUseCase(employeeRepo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{employeeRepo: employeeRepo}
}

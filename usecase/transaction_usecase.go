package usecase

import (
	"fmt"
	"time"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type TransactionUseCase interface {
	RegisterNewTransaction(newData model.Transaction) error
	FindAllTransaction() ([]model.Transaction, error)
	FindTransactionById(id string) (model.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	vehicleUseCase  VehicleUseCase
	customerUseCase CustomerUseCase
	employeeUseCase EmployeeUseCase
}

func (t *transactionUsecase) RegisterNewTransaction(newData model.Transaction) error {
	// get vehicle
	vehicle, err := t.vehicleUseCase.GetVehicle(newData.Vehicle.Id)
	if err != nil {
		return fmt.Errorf("Vehicle with ID: %s not exists", newData.Vehicle.Id)
	}

	// gte customer
	customer, err := t.customerUseCase.GetCustomer(newData.Customer.Id)
	if err != nil {
		return fmt.Errorf("Customer with ID: %s not exists", newData.Customer.Id)
	}

	// get employee
	employee, err := t.employeeUseCase.GetEmployee(newData.Employee.Id)
	if err != nil {
		return fmt.Errorf("Employee with ID: %s not exists", newData.Employee.Id)
	}

	newData.Vehicle = vehicle
	newData.Customer = customer
	newData.Employee = employee
	td, _ := time.Parse("2006-01-02", time.Now().String())
	newData.TransactionDate = td

	return t.transactionRepo.Create(newData)
}
func (t *transactionUsecase) FindAllTransaction() ([]model.Transaction, error) {
	return t.transactionRepo.List()
}
func (t *transactionUsecase) FindTransactionById(id string) (model.Transaction, error) {
	return t.transactionRepo.Get(id)
}

func NewTransactionUseCase(
	transactionRepo repository.TransactionRepository,
	vehicleUseCase VehicleUseCase,
	customerUseCase CustomerUseCase,
	employeeUseCase EmployeeUseCase) TransactionUseCase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		vehicleUseCase:  vehicleUseCase,
		customerUseCase: customerUseCase,
		employeeUseCase: employeeUseCase,
	}
}

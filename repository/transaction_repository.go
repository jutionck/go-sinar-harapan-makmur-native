package repository

import (
	"database/sql"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type TransactionRepository interface {
	Create(newData entity.Transaction) error
	List() ([]dto.TransactionResponseDto, error)
	GetAll() ([]entity.Transaction, error)
	Get(id string) (entity.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(newData entity.Transaction) error {
	sql := "INSERT INTO transaction (id, transaction_date, vehicle_id, customer_id, employee_id, type, qty, payment_amount) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err := t.db.Exec(sql, newData.Id, newData.TransactionDate, newData.Vehicle.Id, newData.Customer.Id, newData.Employee.Id, newData.Type, newData.Qty, newData.PaymentAmount)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) List() ([]dto.TransactionResponseDto, error) {
	sql := `
	select 
	t.id,
	t.transaction_date, 
	c.id as customer_id,
	c.first_name || ' ' || c.last_name,
	v.id as vehicle_id,
	v.brand,
	v.model,
	e.id as employee_id,
	e.first_name || ' ' || e.last_name,
	t.qty,
	t.type,
	t.payment_amount
from
	transaction t
inner join vehicle v on v.id = t.vehicle_id 
inner join customer c on c.id = t.customer_id 
inner join employee e on e.id = t.employee_id`

	var transactions []dto.TransactionResponseDto
	rows, err := t.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction dto.TransactionResponseDto
		err := rows.Scan(
			&transaction.Id,
			&transaction.TransactionDate,
			&transaction.CustomerId,
			&transaction.CustomerName,
			&transaction.VehicleId,
			&transaction.VehicleBrand,
			&transaction.VehicleModel,
			&transaction.EmployeeId,
			&transaction.EmployeeName,
			&transaction.Qty,
			&transaction.Type,
			&transaction.PaymentAmount,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t *transactionRepository) Get(id string) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

func (t *transactionRepository) GetAll() ([]entity.Transaction, error) {
	sql := "SELECT id,transaction_date,customer_id,vehicle_id,employee_id, qty, type, payment_amount FROM transaction ORDER BY transaction_date ASC"

	var transactions []entity.Transaction
	rows, err := t.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.TransactionDate,
			&transaction.Customer.Id,
			&transaction.Vehicle.Id,
			&transaction.Employee.Id,
			&transaction.Qty,
			&transaction.Type,
			&transaction.PaymentAmount,
		)
		if err != nil {
			return nil, err
		}
		sqlVehicle := "SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle WHERE id = $1"
		row := t.db.QueryRow(sqlVehicle, transaction.Vehicle.Id)
		var vehicle entity.Vehicle
		err = row.Scan(&vehicle.Id, &vehicle.Brand, &vehicle.Model, &vehicle.ProductionYear, &vehicle.Color, &vehicle.IsAutomatic, &vehicle.SalePrice, &vehicle.Stock, &vehicle.Status)
		if err != nil {
			return nil, err
		}
		transaction.Vehicle = vehicle

		sqlCustomer := "SELECT id, first_name, last_name, address, email, phone_number, bod FROM customer WHERE id = $1"
		row = t.db.QueryRow(sqlCustomer, transaction.Customer.Id)
		var customer entity.Customer
		err = row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address, &customer.Email, &customer.PhoneNumber, &customer.Bod)
		if err != nil {
			return nil, err
		}
		transaction.Customer = customer

		sqlEmployee := "SELECT id, first_name, last_name, address, email, phone_number, bod, salary, position FROM employee WHERE id = $1"
		row = t.db.QueryRow(sqlEmployee, transaction.Employee.Id)
		var employee entity.Employee
		err = row.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Address, &employee.Email, &employee.PhoneNumber, &employee.Bod, &employee.Salary, &employee.Position)
		if err != nil {
			return nil, err
		}

		sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
		row = t.db.QueryRow(sqlManager, employee.Id)
		var manager entity.Employee
		err = row.Scan(&manager.Id, &manager.FirstName, &manager.LastName, &manager.Address, &manager.Email, &manager.PhoneNumber, &manager.Bod, &manager.Salary, &manager.Position)
		if err != nil {
			return nil, err
		}
		manager.Manager = nil // set manager menjadi nil agar tidak terjadi infinite loop dalam reference circular
		employee.Manager = &manager
		transaction.Employee = employee
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

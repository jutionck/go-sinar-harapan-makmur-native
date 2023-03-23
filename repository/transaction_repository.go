package repository

import (
	"database/sql"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type TransactionRepository interface {
	Create(newData entity.Transaction) error
	List() ([]entity.Transaction, error)
	Get(id string) (entity.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(newData entity.Transaction) error {
	sql := "INSERT INTO transaction (id, transaction_date, vehicle_id, customer_id, employee_id, type, payment_amount) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	_, err := t.db.Exec(sql, newData.Id, newData.TransactionDate, newData.Vehicle.Id, newData.Customer.Id, newData.Employee.Id, newData.Type, newData.PaymentAmount)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) List() ([]entity.Transaction, error) {
	return nil, nil
}
func (t *transactionRepository) Get(id string) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

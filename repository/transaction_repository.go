package repository

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type TransactionRepository interface {
	Create(newData model.Transaction) error
	List() ([]model.Transaction, error)
	Get(id string) (model.Transaction, error)
}

type transactionRepository struct {
	db []model.Transaction
}

func (t *transactionRepository) Create(newData model.Transaction) error {
	t.db = append(t.db, newData)
	if len(t.db) == 0 {
		return fmt.Errorf("Gagal menyimpan data")
	}
	return nil
}

func (t *transactionRepository) List() ([]model.Transaction, error) {
	if len(t.db) == 0 {
		return nil, fmt.Errorf("Database kosong")
	}
	return t.db, nil
}
func (t *transactionRepository) Get(id string) (model.Transaction, error) {
	for _, transaction := range t.db {
		if transaction.Id == id {
			return transaction, nil
		}
	}
	return model.Transaction{}, fmt.Errorf("Data tidak ditemukan")
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              string
	TransactionDate time.Time
	Vehicle
	Customer
	Employee
	Type          string // enum: "Online" & "Offline"
	Qty           int
	PaymentAmount int
}

func (t *Transaction) IsValidType() bool {
	return t.Type == "Online" || t.Type == "Offline"
}

func (t *Transaction) SetId() {
	t.Id = uuid.New().String()
}

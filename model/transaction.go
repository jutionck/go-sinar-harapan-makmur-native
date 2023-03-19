package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID
	TransactionDate time.Time
	Vehicle
	Customer
	Employee
	Type          string // enum: "Online" & "Offline"
	PaymentAmount int
}

func (t *Transaction) IsValidType() bool {
	return t.Type == "Online" || t.Type == "Offline"
}

func (t *Transaction) SetId() {
	t.Id = uuid.New()
}

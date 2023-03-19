package model

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	Id          string
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Email       string
	Bod         time.Time
	Posisition  string
	Salary      int
	Manager     *Employee
}

func (e *Employee) SetId() {
	e.Id = uuid.New().String()
}

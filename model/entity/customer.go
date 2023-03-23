package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          string
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Email       string
	Bod         time.Time
}

func (c *Customer) SetId() {
	c.Id = uuid.New().String()
}

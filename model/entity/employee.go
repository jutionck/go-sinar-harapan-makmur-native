package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type Employee struct {
	Id          sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	Address     sql.NullString
	PhoneNumber sql.NullString
	Email       sql.NullString
	Bod         sql.NullTime
	Position    sql.NullString
	Salary      sql.NullInt64
	Manager     *Employee
}

func (e *Employee) SetId() {
	e.Id = sql.NullString{String: uuid.New().String()}
}

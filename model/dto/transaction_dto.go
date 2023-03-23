package dto

import "time"

type TransactionResponseDto struct {
	Id              string
	TransactionDate time.Time
	VehicleId       string
	VehicleBrand    string
	VehicleModel    string
	CustomerId      string
	CustomerName    string
	EmployeeId      string
	EmployeeName    string
	Type            string
	PaymentAmount   int
}

package dto

import "time"

type TransactionResponseDto struct {
	Id              string
	TransactionDate time.Time
	CustomerId      string
	CustomerName    string
	VehicleId       string
	VehicleBrand    string
	VehicleModel    string
	EmployeeId      string
	EmployeeName    string
	Qty             int
	Type            string
	PaymentAmount   int
}

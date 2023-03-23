package entity

import "github.com/google/uuid"

type Vehicle struct {
	Id             string
	Brand          string
	Model          string
	ProductionYear int
	Color          string
	IsAutomatic    bool
	Stock          int
	SalePrice      int
	Status         string // enum: "Baru" & "Bekas"
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "Baru" || v.Status == "Bekas"
}

func (v *Vehicle) SetId() {
	v.Id = uuid.New().String()
}

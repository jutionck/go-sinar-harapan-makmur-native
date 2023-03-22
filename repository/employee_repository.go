package repository

import (
	"database/sql"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	GetByEmail(email string) (model.Employee, error)
	GetByPhoneNumber(phoneNumber string) (model.Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(newData model.Employee) error {
	sql := "INSERT INTO employee (id, first_name, last_name, address, phone_number, email, bod, position, salary, manager_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}

	_, err := e.db.Exec(sql, newData.Id, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod, newData.Posisition, newData.Salary, managerID)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) List() ([]model.Employee, error) {
	// sementara (manager_id di hilangkan)
	// harusnya join ini untuk mendapatkan managerID
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee`
	rows, err := e.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Address, &employee.PhoneNumber, &employee.Email, &employee.Bod, &employee.Posisition, &employee.Salary)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *employeeRepository) Get(id string) (model.Employee, error) {
	// sementara (manager_id di hilangkan)
	// harusnya join ini untuk mendapatkan managerID
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee WHERE id = $1`
	var employee model.Employee
	err := e.db.QueryRow(sql, id).Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Address, &employee.PhoneNumber, &employee.Email, &employee.Bod, &employee.Posisition, &employee.Salary)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) Update(newData model.Employee) error {
	sql := "UPDATE employee SET first_name = $1, last_name = $2, address = $3, phone_number = $4, email = $5, bod = $6, position = $7, salary = $8, manager_id = $9 WHERE id = $10"
	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}
	_, err := e.db.Exec(sql, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod, newData.Posisition, newData.Salary, managerID, newData.Id)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) Delete(id string) error {
	sql := "DELETE FROM employee WHERE id = $1"
	_, err := e.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) GetByEmail(email string) (model.Employee, error) {
	sql := `SELECT id, email FROM employee WHERE email = $1`
	var employee model.Employee
	err := e.db.QueryRow(sql, email).Scan(&employee.Id, &employee.Email)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetByPhoneNumber(phoneNumber string) (model.Employee, error) {
	sql := `SELECT id, phone_number FROM employee WHERE phone_number = $1`
	var employee model.Employee
	err := e.db.QueryRow(sql, phoneNumber).Scan(&employee.Id, &employee.PhoneNumber)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

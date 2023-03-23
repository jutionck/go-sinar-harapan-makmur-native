package repository

import (
	"database/sql"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type EmployeeRepository interface {
	BaseRepository[entity.Employee]
	GetByEmail(email string) (entity.Employee, error)
	GetByPhoneNumber(phoneNumber string) (entity.Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(newData entity.Employee) error {
	sql := "INSERT INTO employee (id, first_name, last_name, address, phone_number, email, bod, position, salary, manager_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}

	_, err := e.db.Exec(sql, newData.Id, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod, newData.Position, newData.Salary, managerID)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) List() ([]entity.Employee, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee`
	rows, err := e.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var employees []entity.Employee
	for rows.Next() {
		var employee entity.Employee
		err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Address, &employee.PhoneNumber, &employee.Email, &employee.Bod, &employee.Position, &employee.Salary)
		if err != nil {
			return nil, err
		}

		sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
		row := e.db.QueryRow(sqlManager, employee.Id)
		var manager entity.Employee
		err = row.Scan(&manager.Id, &manager.FirstName, &manager.LastName, &manager.Address, &manager.Email, &manager.PhoneNumber, &manager.Bod, &manager.Salary, &manager.Position)
		if err != nil {
			return nil, err
		}
		manager.Manager = nil // set manager menjadi nil agar tidak terjadi infinite loop dalam reference circular
		employee.Manager = &manager
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *employeeRepository) Get(id string) (entity.Employee, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee WHERE id = $1`
	var employee entity.Employee
	err := e.db.QueryRow(sql, id).Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Address, &employee.PhoneNumber, &employee.Email, &employee.Bod, &employee.Position, &employee.Salary)
	if err != nil {
		return entity.Employee{}, err
	}
	sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
	row := e.db.QueryRow(sqlManager, employee.Id)
	var manager entity.Employee
	err = row.Scan(&manager.Id, &manager.FirstName, &manager.LastName, &manager.Address, &manager.Email, &manager.PhoneNumber, &manager.Bod, &manager.Salary, &manager.Position)
	if err != nil {
		return entity.Employee{}, err
	}
	employee.Manager = &manager

	return employee, nil
}

func (e *employeeRepository) Update(newData entity.Employee) error {
	sql := "UPDATE employee SET first_name = $1, last_name = $2, address = $3, phone_number = $4, email = $5, bod = $6, position = $7, salary = $8, manager_id = $9 WHERE id = $10"
	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}
	_, err := e.db.Exec(sql, newData.FirstName, newData.LastName, newData.Address, newData.PhoneNumber, newData.Email, newData.Bod, newData.Position, newData.Salary, managerID, newData.Id)
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

func (e *employeeRepository) GetByEmail(email string) (entity.Employee, error) {
	sql := `SELECT id, email FROM employee WHERE email = $1`
	var employee entity.Employee
	err := e.db.QueryRow(sql, email).Scan(&employee.Id, &employee.Email)
	if err != nil {
		return entity.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetByPhoneNumber(phoneNumber string) (entity.Employee, error) {
	sql := `SELECT id, phone_number FROM employee WHERE phone_number = $1`
	var employee entity.Employee
	err := e.db.QueryRow(sql, phoneNumber).Scan(&employee.Id, &employee.PhoneNumber)
	if err != nil {
		return entity.Employee{}, err
	}
	return employee, nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

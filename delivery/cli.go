package delivery

import (
	"fmt"
	"time"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
)

func VehicleCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	vehicleRepository := repository.NewVehicleRepository(dbConn.Conn())
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepository)

	// cretae
	newVehicle := model.Vehicle{
		Brand:     "Toyota",
		Model:     "Alphard",
		Color:     "Putih",
		Stock:     0,
		Status:    "Baru",
		SalePrice: 900000000,
	}
	newVehicle.SetId()
	if err := vehicleUseCase.RegisterNewVehicle(newVehicle); err != nil {
		fmt.Println(err)
		return
	}

	// Get All
	vehicles, err := vehicleUseCase.FindAllVehicle()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range vehicles {
			fmt.Println("ID:", v.Id)
			fmt.Println("Brand:", v.Brand)
			fmt.Println("Model:", v.Model)
			fmt.Println("Production Year:", v.ProductionYear)
			fmt.Println("Color:", v.Color)
			fmt.Println("Price:", v.SalePrice)
			fmt.Println("Is Automatic:", v.IsAutomatic)
			fmt.Println("Stock:", v.Stock)
			fmt.Println("Status:", v.Status)
			fmt.Println()
		}
	}
}

func CustomerCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	customerRepository := repository.NewCustomerRepository(dbConn.Conn())
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)

	// cretae
	bod, _ := time.Parse("2006-01-02", "1999-11-11")
	newCustomer := model.Customer{
		FirstName:   "Tika",
		LastName:    "Yesi",
		PhoneNumber: "0821444444",
		Email:       "tika.yesi@gmail.com",
		Bod:         bod,
	}
	newCustomer.SetId()
	if err := customerUseCase.RegisterNewCustomer(newCustomer); err != nil {
		fmt.Println(err)
		return
	}

	// Get All
	customers, err := customerUseCase.FindAllCustomer()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, c := range customers {
			fmt.Println("ID:", c.Id)
			fmt.Println("Name:", c.FirstName, c.LastName)
			fmt.Println("Phone Number:", c.PhoneNumber)
			fmt.Println("Email:", c.Email)
			fmt.Println("Birth Date:", c.Bod)
			fmt.Println()
		}
	}
}

func EmployeeCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	employeeRepository := repository.NewEmployeeRepository(dbConn.Conn())
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository)

	// cretae
	bod, _ := time.Parse("2006-01-02", "1990-11-11")
	newEmployee := model.Employee{
		FirstName:   "Tika",
		LastName:    "Yesi",
		PhoneNumber: "0821444444",
		Email:       "tika.yesi@gmail.com",
		Bod:         bod,
		Posisition:  "Software Developer",
		Salary:      15000000,
		Manager:     &model.Employee{Id: "34258ecc-b35c-4da9-8574-c452475af11f"},
	}
	newEmployee.SetId()
	if err := employeeUseCase.RegisterNewEmployee(newEmployee); err != nil {
		fmt.Println(err)
		return
	}

	// Get All
	employees, err := employeeUseCase.FindAllEmployee()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, e := range employees {
			fmt.Println("ID:", e.Id)
			fmt.Println("Name:", e.FirstName, e.LastName)
			fmt.Println("Phone Number:", e.PhoneNumber)
			fmt.Println("Email:", e.Email)
			fmt.Println("Birth Date:", e.Bod)
			fmt.Println("Position:", e.Posisition)
			fmt.Println("Salary:", e.Salary)
			fmt.Println("Manager:", e.Manager)
			fmt.Println()
		}
	}
}

// func TransactionCLI() {
// 	transactionRepo := repository.NewTransactionRepository()
// 	vehicleUseCase := usecase.NewVehicleUseCase(repository.NewVehicleRepository())
// 	customerUseCase := usecase.NewCustomerUseCase(repository.NewCustomerRepository())
// 	employeeUseCase := usecase.NewEmployeeUseCase(repository.NewEmployeeRepository())

// 	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo, vehicleUseCase, customerUseCase, employeeUseCase)
// 	newTransaction := model.Transaction{
// 		Id:            "T0001",
// 		Vehicle:       model.Vehicle{Id: "V0001"},
// 		Customer:      model.Customer{Id: "C0001"},
// 		Employee:      model.Employee{Id: "EP001"},
// 		Type:          "Online",
// 		PaymentAmount: 301000000,
// 	}

// 	if err := transactionUseCase.RegisterNewTransaction(newTransaction); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Get All
// 	transactions, err := transactionUseCase.FindAllTransaction()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		for _, t := range transactions {
// 			fmt.Println("ID:", t.Id)
// 			fmt.Println("Date:", t.TransactionDate)
// 			fmt.Println("Vehicle:", t.Vehicle.Brand, t.Vehicle.Model)
// 			fmt.Println("Customer:", t.Customer.FirstName, t.Customer.LastName)
// 			fmt.Println("Employee:", t.Employee.FirstName, t.Employee.LastName)
// 			fmt.Println("Type:", t.Type)
// 			fmt.Println("Payment Amount:", t.PaymentAmount)
// 			fmt.Println()
// 		}
// 	}
// }

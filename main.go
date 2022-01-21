package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	EmployeeId   int
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	HireDate     string
	JobId        int
	Salary       float32
	ManagerId    int
	DepartmentId int
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "")

	if err != nil {
		panic(err)
	}

	// Add Employee
	// employee := Employee{FirstName: "Joe", LastName: "K", Email: "", HireDate: time.Now().Format("2006-01-02"), JobId: 1, Salary: 8000}
	// err = AddEmployee(employee)
	// if err != nil {
	// 	panic(err)
	// }

	// Update Employee
	// employee := Employee{Email: "", Salary: 10000, EmployeeId: 207}
	// err = UpdateEmployee(employee)
	// if err != nil {
	// 	panic(err)
	// }

	// Delete Employee
	// err = DeleteEmployee(207)
	// if err != nil {
	// 	panic(err)
	// }

	employees, err := GetEmployees()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, employee := range employees {
		fmt.Println(employee)
	}

	employee, err := GetEmployee(116)

	if err != nil {
		panic(err)
	}

	fmt.Println(employee)

}

func GetEmployees() ([]Employee, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT employee_id, first_name, last_name FROM employees"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.EmployeeId, &employee.FirstName, &employee.LastName)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func GetEmployee(id int) (*Employee, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT employee_id, first_name, last_name FROM employees WHERE employee_id=?"

	row := db.QueryRow(query, id)

	employee := Employee{}

	err = row.Scan(&employee.EmployeeId, &employee.FirstName, &employee.LastName)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func AddEmployee(employee Employee) error {

	query := "INSERT INTO employees(first_name, last_name, email, hire_date, job_id, salary) values(?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, employee.FirstName, employee.LastName, employee.Email, employee.HireDate, employee.JobId, employee.Salary)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return nil
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil
}

func UpdateEmployee(employee Employee) error {

	query := "UPDATE employees SET email=?, salary=? WHERE employee_id=? "

	result, err := db.Exec(query, employee.Email, employee.Salary, employee.EmployeeId)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return nil
	}

	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil
}

func DeleteEmployee(id int) error {

	query := "DELETE FROM employees WHERE employee_id=? "

	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return nil
	}

	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}

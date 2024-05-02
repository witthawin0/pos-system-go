package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/witthawin0/pos-system-go/internal/domain"
)

// ProductRepositoryImpl implements the ProductRepository interface.
type employeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeReposistoryImpl(db *sql.DB) *employeeRepositoryImpl {
	return &employeeRepositoryImpl{db: db}
}

func (r *employeeRepositoryImpl) AddEmployee(employee domain.Employee) (int, error) {
	// Prepare the SQL query
	query := `INSERT INTO employees (firstname, lastname, date_of_birth, username, password, role) VALUES (?, ?, ?, ?, ?, ?)`

	// Execute the query with the order data
	result, err := r.db.Exec(query, employee.FirstName, employee.LastName, employee.DateOfBirth, employee.Username, employee.Password, employee.Role)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

func (r *employeeRepositoryImpl) UpdateEmployee(id int, employee domain.Employee) error {
	// Prepare the SQL query
	query := `UPDATE employees SET firstname = ?, lastname = ?, date_of_birth = ?, username =?, password = ?, role = ? WHERE id = ?`

	// Execute the query with the order data
	_, err := r.db.Exec(query, employee.FirstName, employee.LastName, employee.DateOfBirth, employee.Username, employee.Password, employee.Role, id)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func (r *employeeRepositoryImpl) RemoveEmployee(id int) error {
	// Prepare the SQL query
	query := `DELETE FROM employees WHERE id = ?`

	// Execute the query with the order ID
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (r *employeeRepositoryImpl) GetEmployeeByID(id int) (*domain.Employee, error) {
	var employee domain.Employee

	// Prepare the SQL query
	query := `SELECT id, firstname, lastname, date_of_birth, username, password, role FROM employees WHERE id = ?`

	// Execute the query
	row := r.db.QueryRow(query, id)

	// Scan the result into an Order struct
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.DateOfBirth, &employee.Username, &employee.Password, &employee.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where order with given ID is not found
			return &domain.Employee{}, errors.New("employee not found")
		}
		// Handle other errors
		return &domain.Employee{}, err
	}

	return &employee, nil

}

func (r *employeeRepositoryImpl) ListEmployees() ([]*domain.Employee, error) {
	// Prepare the SQL query
	query := `SELECT id, firstname, lastname, date_of_birth, username, password, role, created_at, updated_at FROM employees`

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and scan each row into an Order struct
	var employees []*domain.Employee
	for rows.Next() {
		var emp domain.Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.DateOfBirth, &emp.Username, &emp.Password, &emp.Role, &emp.CreatedAt, &emp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &emp)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

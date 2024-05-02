package dto

import "time"

// EmployeeAddDTO represents the data required to add a new employee.
type EmployeeAddDTO struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}

// EmployeeUpdateDTO represents the data required to update an existing employee.
type EmployeeUpdateDTO struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}

// EmployeeDTO represents the data for a single employee.
type EmployeeDTO struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// EmployeeListDTO represents a list of employees.
type EmployeeListDTO struct {
	Employees []*EmployeeDTO `json:"employees"`
}

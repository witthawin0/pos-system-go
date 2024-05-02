package domain

import "time"

// Product represents a product in the store.
type Product struct {
	ID         int
	Name       string
	Price      float64
	Quantity   int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Order represents an order placed by a customer.
type Order struct {
	ID         int
	CustomerID int
	Products   []Product
	TotalPrice float64
	Status     string
}

// Customer represents a customer.
type Customer struct {
	ID    int
	Name  string
	Email string
}

// Employee represents a Employee.
type Employee struct {
	ID          int
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Username    string
	Password    string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

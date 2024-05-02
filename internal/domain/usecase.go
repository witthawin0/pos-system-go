package domain

import "github.com/witthawin0/pos-system-go/internal/dto"

// OrderUseCase defines the interface for order-related use cases.
type OrderUseCase interface {
	CreateOrder(customerID int, products []Product) (Order, error)
	CalculateTotalPrice(products []Product) (float64, error)
}

// InventoryUseCase defines the interface for inventory-related use cases.
type InventoryUseCase interface {
	AddProduct(name string, price float64, quantity int, categoryID int) (int, error)
	UpdateProduct(id int, name string, price float64, quantity int, categoryID int) error
	RemoveProduct(id int) error
	GetProductByID(id int) (Product, error)
	ListProducts() ([]Product, error)
	IncreaseStock(id int, quantity int) error
	DecreaseStock(id int, quantity int) error
	SetStockLevel(id int, quantity int) error
}

// EmployeeUseCase defines the interface for employee-related use cases.
type EmployeeUseCase interface {
	AddEmployee(employee dto.EmployeeAddDTO) (int, error)
	UpdateEmployee(id int, employee dto.EmployeeUpdateDTO) error
	RemoveEmployee(id int) error
	GetEmployeeByID(id int) (*dto.EmployeeDTO, error)
	ListEmployees() ([]dto.EmployeeDTO, error)
}

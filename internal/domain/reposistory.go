package domain

// ProductRepository defines the interface for product-related operations.
type ProductRepository interface {
	GetProductByID(id int) (*Product, error)
	GetProducts() ([]*Product, error)
	SaveProduct(product Product) error
	UpdateProduct(product Product) error
	DeleteProduct(id int) error
}

// OrderRepository defines the interface for order-related operations.
type OrderRepository interface {
	GetOrderByID(id int) (*Order, error)
	GetOrdersByCustomerID(customerID int) ([]*Order, error)
	SaveOrder(order Order) error
	UpdateOrder(order Order) error
	DeleteOrder(id int) error
}

// EmployeeRepository defines the interface for Employee-related operations.
type EmployeeRepository interface {
	AddEmployee(employee Employee) (int, error)
	UpdateEmployee(id int, employee Employee) error
	RemoveEmployee(id int) error
	GetEmployeeByID(id int) (*Employee, error)
	ListEmployees() ([]*Employee, error)
}

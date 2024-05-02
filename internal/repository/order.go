package repository

import (
	"database/sql"
	"errors"

	"github.com/witthawin0/pos-system-go/internal/domain"
)

// OrderRepositoryImpl implements the OrderRepository interface.
type orderRepositoryImpl struct {
	db *sql.DB
}

// NewOrderRepositoryImpl creates a new instance of OrderRepositoryImpl.
func NewOrderRepositoryImpl(db *sql.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) GetOrderByID(id int) (domain.Order, error) {
	var order domain.Order

	// Prepare the SQL query
	query := "SELECT id, customer_id, total_price, status FROM orders WHERE id = ?"

	// Execute the query
	row := r.db.QueryRow(query, id)

	// Scan the result into an Order struct
	err := row.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where order with given ID is not found
			return domain.Order{}, errors.New("order not found")
		}
		// Handle other errors
		return domain.Order{}, err
	}

	return order, nil
}

func (r *orderRepositoryImpl) GetOrdersByCustomerID(customerID int) ([]domain.Order, error) {
	var orders []domain.Order

	// Prepare the SQL query
	query := "SELECT id, customer_id, total_price, status FROM orders WHERE customer_id = ?"

	// Execute the query
	rows, err := r.db.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and scan each row into an Order struct
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepositoryImpl) SaveOrder(order domain.Order) error {
	// Prepare the SQL query
	query := "INSERT INTO orders (customer_id, total_price, status) VALUES (?, ?, ?)"

	// Execute the query with the order data
	result, err := r.db.Exec(query, order.CustomerID, order.TotalPrice, order.Status)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted order
	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Update the order ID in the Order struct
	order.ID = int(orderID)

	return nil
}

func (r *orderRepositoryImpl) UpdateOrder(order domain.Order) error {
	// Prepare the SQL query
	query := "UPDATE orders SET customer_id = ?, total_price = ?, status = ? WHERE id = ?"

	// Execute the query with the order data
	_, err := r.db.Exec(query, order.CustomerID, order.TotalPrice, order.Status, order.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepositoryImpl) DeleteOrder(id int) error {
	// Prepare the SQL query
	query := "DELETE FROM orders WHERE id = ?"

	// Execute the query with the order ID
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

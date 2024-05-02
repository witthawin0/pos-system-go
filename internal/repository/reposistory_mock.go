package repository

import "github.com/witthawin0/pos-system-go/internal/domain"

type MockProductRepository struct{}

func (m *MockProductRepository) GetProductByID(id int) (domain.Product, error) {
	// Mock implementation, return a dummy product
	return domain.Product{
		ID:       id,
		Name:     "Mock Product",
		Price:    10.0,
		Quantity: 100,
	}, nil
}

type MockOrderRepository struct{}

func (m *MockOrderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	// Mock implementation, return the same order
	return order, nil
}

package usecase

import (
	"errors"

	"github.com/witthawin0/pos-system-go/internal/domain"
)

// OrderUseCaseImpl implements the OrderUseCase interface.
type OrderUseCaseImpl struct {
	productRepo domain.ProductRepository
	orderRepo   domain.OrderRepository
}

// NewOrderUseCaseImpl creates a new instance of OrderUseCaseImpl.
func NewOrderUseCaseImpl(productRepo domain.ProductRepository, orderRepo domain.OrderRepository) *OrderUseCaseImpl {
	return &OrderUseCaseImpl{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (uc *OrderUseCaseImpl) CreateOrder(customerID int, products []domain.Product) (domain.Order, error) {
	// Validate customer ID
	if customerID <= 0 {
		return domain.Order{}, errors.New("invalid customer ID")
	}

	// Validate products
	if len(products) == 0 {
		return domain.Order{}, errors.New("no products specified")
	}

	// Calculate total price
	totalPrice, err := uc.CalculateTotalPrice(products)
	if err != nil {
		return domain.Order{}, err
	}

	// Create order
	order := domain.Order{
		CustomerID: customerID,
		Products:   products,
		TotalPrice: totalPrice,
		Status:     "pending",
	}

	// Save order
	err = uc.orderRepo.SaveOrder(order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderUseCaseImpl) CalculateTotalPrice(products []domain.Product) (float64, error) {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.Price * float64(product.Quantity)
	}
	return totalPrice, nil
}

// InventoryUseCaseImpl implements the InventoryUseCase interface.
type InventoryUseCaseImpl struct {
	productRepo domain.ProductRepository
}

// NewInventoryUseCaseImpl creates a new instance of InventoryUseCaseImpl.
func NewInventoryUseCaseImpl(productRepo domain.ProductRepository) *InventoryUseCaseImpl {
	return &InventoryUseCaseImpl{
		productRepo: productRepo,
	}
}

func (uc *InventoryUseCaseImpl) UpdateInventory(products []domain.Product) error {
	for _, product := range products {
		// Update inventory for each product
		err := uc.productRepo.UpdateProduct(product)
		if err != nil {
			return err
		}
	}
	return nil
}

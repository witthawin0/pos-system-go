package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/witthawin0/pos-system-go/internal/domain"
)

type inventoryHandler struct {
	inventoryUseCase domain.InventoryUseCase
}

func NewInventoryHandler(inventoryUseCase domain.InventoryUseCase) inventoryHandler {
	return inventoryHandler{inventoryUseCase: inventoryUseCase}
}

func (h *employeeHandler) AddProduct(c echo.Context) error {
	return nil
}

func (h *employeeHandler) UpdateProduct(c echo.Context) error {
	return nil
}

func (h *employeeHandler) RemoveProduct(c echo.Context) error {
	return nil
}

func (h *employeeHandler) GetProductByID(c echo.Context) error {
	return nil
}

func (h *employeeHandler) ListProducts(c echo.Context) error {
	return nil
}

func (h *employeeHandler) IncreaseStock(c echo.Context) error {
	return nil
}

func (h *employeeHandler) DecreaseStock(c echo.Context) error {
	return nil
}

func (h *employeeHandler) SetStockLevel(c echo.Context) error {
	return nil
}

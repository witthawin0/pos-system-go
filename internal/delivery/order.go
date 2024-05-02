package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/witthawin0/pos-system-go/internal/domain"
)

type orderHandler struct {
	orderUseCase domain.OrderUseCase
}

func NewOrderHandler(orderUseCase domain.OrderUseCase) orderHandler {
	return orderHandler{orderUseCase: orderUseCase}
}

// func (h *OrderHandler) GetOrderHandler(c echo.Context) error {
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid order ID")
// 	}

// 	// Use your order use case to get the order
// 	order, err := h.OrderUseCase.(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Order not found")
// 	}

// 	// Marshal the order to JSON and send it in the response
// 	return c.JSON(http.StatusOK, order)
// }

func (h *orderHandler) CreateOrderHandler(c echo.Context) error {
	// Parse the request body into an Order struct
	var requestData struct {
		CustomerID int              `json:"customer_id"`
		Products   []domain.Product `json:"products"`
	}

	err := c.Bind(&requestData)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// Use your order use case to create the order
	createdOrder, err := h.orderUseCase.CreateOrder(requestData.CustomerID, requestData.Products)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create order")
	}

	//send created order in the response
	return c.JSON(http.StatusCreated, createdOrder)
}

// func (h *OrderHandler) UpdateOrderHandler(c echo.Context) error {
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid order ID")
// 	}

// 	// Parse the request body into an Order struct
// 	var order Order
// 	err = json.NewDecoder(r.Body).Decode(&order)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
// 	}

// 	// Use your order use case to update the order
// 	updatedOrder, err := h.orderUseCase.UpdateOrder(id, order)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update order")
// 	}

// 	// Marshal the updated order to JSON and send it in the response
// 	json.NewEncoder(w).Encode(updatedOrder)
// }

// func (h *OrderHandler) DeleteOrderHandler(c echo.Context) error {
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid order ID")
// 	}

// 	// Use your order use case to delete the order
// 	err = orderUseCase.DeleteOrder(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete order")
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

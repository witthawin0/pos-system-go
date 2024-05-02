package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/witthawin0/pos-system-go/internal/domain"
	"github.com/witthawin0/pos-system-go/internal/dto"
)

type employeeHandler struct {
	uc domain.EmployeeUseCase
}

func NewEmployeeHandler(uc domain.EmployeeUseCase) employeeHandler {
	return employeeHandler{uc: uc}
}

func (h *employeeHandler) AddEmployee(c echo.Context) error {
	// Parse the request body into an Employee struct
	var employee dto.EmployeeAddDTO

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call the AddEmployee method from the EmployeeUseCase
	id, err := h.uc.AddEmployee(employee)
	if err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to add employee"})
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": id})

}

func (h *employeeHandler) UpdateEmployee(c echo.Context) error {
	//parse id to int
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request id"})

	}

	var employee dto.EmployeeUpdateDTO

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call the AddEmployee method from the EmployeeUseCase
	err = h.uc.UpdateEmployee(id, employee)
	if err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update employee"})
	}

	return c.JSON(http.StatusBadRequest, map[string]int{"id": id})

}

func (h *employeeHandler) RemoveEmployee(c echo.Context) error {
	//parse id to int
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request id"})

	}

	err = h.uc.RemoveEmployee(id)

	if err != nil {
		return c.JSON(http.StatusNoContent, map[string]string{"error": "failed to update employee"})

	}

	return c.JSON(http.StatusOK, map[string]string{"ok": "employee is updated sucessfully"})

}

func (h *employeeHandler) GetEmployeeByID(c echo.Context) error {
	//parse id to int
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request id"})
	}

	employee, err := h.uc.GetEmployeeByID(id)
	if err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, employee)
}

func (h *employeeHandler) ListEmployees(c echo.Context) error {
	employees, err := h.uc.ListEmployees()
	if err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, employees)
}

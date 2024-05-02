package usecase

import (
	"errors"

	"github.com/witthawin0/pos-system-go/internal/domain"
	"github.com/witthawin0/pos-system-go/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

// OrderUseCaseImpl implements the OrderUseCase interface.
type employeeUseCaseImpl struct {
	repo domain.EmployeeRepository
}

// NewOrderUseCaseImpl creates a new instance of OrderUseCaseImpl.
func NewEmployeeUseCaseImpl(repo domain.EmployeeRepository) *employeeUseCaseImpl {
	return &employeeUseCaseImpl{repo: repo}
}

func (uc *employeeUseCaseImpl) AddEmployee(employee dto.EmployeeAddDTO) (int, error) {
	// Perform validation logic here
	if employee.FirstName == "" || employee.LastName == "" || employee.Username == "" || employee.Password == "" || employee.DateOfBirth.String() == "" || employee.Role == "" {
		return 0, errors.New("missing required fields")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Replace the password with the hashed password
	employee.Password = string(hashedPassword)

	newEmployeeData := domain.Employee{
		FirstName:   employee.FirstName,
		LastName:    employee.LastName,
		DateOfBirth: employee.DateOfBirth,
		Username:    employee.Username,
		Password:    employee.Password,
		Role:        employee.Role,
	}

	// Call the repository method to add the employee
	return uc.repo.AddEmployee(newEmployeeData)
}

func (uc *employeeUseCaseImpl) UpdateEmployee(id int, employee dto.EmployeeUpdateDTO) error {
	// Retrieve the existing employee from the repository
	existingEmployee, err := uc.repo.GetEmployeeByID(id)

	if err != nil {
		return err
	}

	// Update the existing employee's fields with the new values if they are not empty
	if employee.FirstName != "" {
		existingEmployee.FirstName = employee.FirstName
	}
	if employee.LastName != "" {
		existingEmployee.LastName = employee.LastName
	}
	if employee.Username != "" {
		existingEmployee.Username = employee.Username
	}
	if !employee.DateOfBirth.IsZero() {
		existingEmployee.DateOfBirth = employee.DateOfBirth
	}
	if employee.Role != "" {
		existingEmployee.Role = employee.Role
	}

	newEmployeeData := domain.Employee{
		ID:          existingEmployee.ID,
		FirstName:   employee.FirstName,
		LastName:    employee.LastName,
		DateOfBirth: employee.DateOfBirth,
		Username:    employee.Username,
		Role:        employee.Role,
	}

	// Call the repository method to update the employee
	return uc.repo.UpdateEmployee(id, newEmployeeData)
}

func (uc *employeeUseCaseImpl) RemoveEmployee(id int) error {
	// Call the repository method to remove the employee
	return uc.repo.RemoveEmployee(id)

}

func (uc *employeeUseCaseImpl) GetEmployeeByID(id int) (*dto.EmployeeDTO, error) {
	// Call the repository method to retrieve the employee by ID
	employee, err := uc.repo.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}

	// Map employee to EmployeeDTO
	res := dto.EmployeeDTO{
		ID:          employee.ID,
		FirstName:   employee.FirstName,
		LastName:    employee.LastName,
		Username:    employee.Username,
		DateOfBirth: employee.DateOfBirth,
		Role:        employee.Role,
	}

	return &res, nil

}

func (uc *employeeUseCaseImpl) ListEmployees() ([]dto.EmployeeDTO, error) {
	// Call the repository method to list all employees
	ems, err := uc.repo.ListEmployees()
	if err != nil {
		return nil, err
	}

	// Map employees to EmployeeDTO
	employees := []dto.EmployeeDTO{}
	for _, em := range ems {
		employee := dto.EmployeeDTO{
			ID:          em.ID,
			FirstName:   em.FirstName,
			LastName:    em.LastName,
			DateOfBirth: em.DateOfBirth,
			Username:    em.Username,
			Role:        em.Role,
		}
		employees = append(employees, employee)
	}
	return employees, nil

}

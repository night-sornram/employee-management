package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/repository"
)

type handleFiber struct {
	service repository.EmployeeService
}

func NewHandleFiber(service repository.EmployeeService) *handleFiber {
	return &handleFiber{
		service: service,
	}
}

func (h *handleFiber) GetEmployees(c *fiber.Ctx) error {
	Employees, err := h.service.GetEmployees()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(Employees)
}

func (h *handleFiber) GetEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	Employee, err := h.service.GetEmployee(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(Employee)
}

func (h *handleFiber) CreateEmployee(c *fiber.Ctx) error {
	var Employee repository.Employee
	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newEmployee, err := h.service.CreateEmployee(Employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(newEmployee)
}

func (h *handleFiber) UpdateEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var Employee repository.Employee
	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	updateEmployee, err := h.service.UpdateEmployee(id, Employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(updateEmployee)
}

func (h *handleFiber) DeleteEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = h.service.DeleteEmployee(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(204)
}

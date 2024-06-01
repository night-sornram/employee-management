package adapter

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"golang.org/x/crypto/bcrypt"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Employees)
}

func (h *handleFiber) GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	Employee, err := h.service.GetEmployee(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Employee)
}

func (h *handleFiber) CreateEmployee(c *fiber.Ctx) error {
	var Employee repository.Employee

	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(Employee.Password), 14)
	Employee.Password = string(password)
	newEmployee, err := h.service.CreateEmployee(Employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "success",
			"data":    newEmployee,
		},
	)
}

func (h *handleFiber) UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var Employee repository.Employee
	if err := c.BodyParser(&Employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	_, err := h.service.UpdateEmployee(id, Employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *handleFiber) DeleteEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = h.service.DeleteEmployee(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handleFiber) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	Employee, err := h.service.Login(data["id"], data["password"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": Employee.Email,
		"iss":   Employee.EmployeeID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"role":  Employee.Role,
	})

	secretKey := os.Getenv("SECRET")

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func (h *handleFiber) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *handleFiber) GetMe(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")
	secretKey := os.Getenv("SECRET")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	token, err := jwt.ParseWithClaims(tokenArr[1], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.MapClaims)

	var Employee repository.Employee

	Employee, err = h.service.GetMe((*claims)["iss"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(Employee)
}

func (h *handleFiber) ChangePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	_, err := h.service.ChangePassword(data["id"], data["password"], data["new_password"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

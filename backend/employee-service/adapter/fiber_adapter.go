package adapter

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/night-sornram/employee-management/employee-service/repository"
)

type handleFiber struct {
	service repository.EmployeeService
}

func NewHandlerFiber(service repository.EmployeeService) *handleFiber {
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

	// -- bcrypt --
	//password, _ := bcrypt.GenerateFromPassword([]byte(Employee.Password), 14)
	//Employee.Password = string(password)

	// -- sha256 --
	password := sha256.Sum256([]byte(Employee.Password))
	Employee.Password = base64.StdEncoding.EncodeToString(password[:])

	// data, _ := json.Marshal(Employee)

	// h.message.PublishOnQueue(data, "employee")

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

func (f *handleFiber) DownloadCSV(c *fiber.Ctx) error {
	query := c.Query("query")
	//if query == "" {
	//	return c.Status(fiber.StatusBadRequest).SendString("Query is missing")
	//}
	data, err := f.service.DownloadCSV(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Nope")
	}
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=data.csv")
	c.Set(fiber.HeaderContentType, "text/csv")
	return c.Send(data)
}

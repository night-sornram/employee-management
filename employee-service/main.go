package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/adapter"
	"github.com/night-sornram/employee-management/middleware"
	"github.com/night-sornram/employee-management/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "employee"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := adapter.NewGormAdapter(db)
	service := repository.NewEmployeeService(repo)
	handle := adapter.NewHandleFiber(service)

	db.AutoMigrate(&repository.Employee{})
	app.Use(cors.New())

	app.Post("/login", handle.Login)
	app.Post("/logout", handle.Logout)
	app.Get("/me", handle.GetMe)
	app.Use("/api", middleware.Protected())
	app.Post("/api/changePassword", handle.ChangePassword)
	app.Get("/api/employees", middleware.Authorize("admin"), handle.GetEmployees)
	app.Get("/api/employees/:id", middleware.Authorize("admin", "user"), handle.GetEmployee)
	app.Post("/api/employees", middleware.Authorize("admin"), handle.CreateEmployee)
	app.Put("/api/employees/:id", middleware.Authorize("admin", "user"), handle.UpdateEmployee)
	app.Delete("/api/employees/:id", middleware.Authorize("admin"), handle.DeleteEmployee)

	app.Listen(":8080")
}

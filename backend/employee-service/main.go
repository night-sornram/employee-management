package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/common-utils"
	"github.com/night-sornram/employee-management/common-utils/middleware"
	"github.com/night-sornram/employee-management/leave-management-service/adapter"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()
	dsn := common_utils.ConnectDB("8080")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := adapter.NewGormAdapter(db)
	service := repository.NewEmployeeService(repo)
	handle := adapter.NewHandleFiber(service)

	err = db.AutoMigrate(&repository.Employee{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	app.Use(cors.New())

	app.Use("/api/employees", middleware.Protected())
	app.Use("/me", middleware.Protected())
	app.Use("/api/changePassword", middleware.Protected())

	app.Post("/login", handle.Login)
	app.Post("/logout", handle.Logout)
	app.Get("/me", handle.GetMe)

	app.Get("/api/employees", middleware.Authorize("admin"), handle.GetEmployees)
	app.Get("/api/employees/:id", middleware.Authorize("admin", "user"), handle.GetEmployee)
	app.Post("/api/employees", middleware.Authorize("admin"), handle.CreateEmployee)
	app.Put("/api/employees/:id", middleware.Authorize("admin", "user"), handle.UpdateEmployee)
	app.Delete("/api/employees/:id", middleware.Authorize("admin"), handle.DeleteEmployee)
	app.Post("/api/changePassword", handle.ChangePassword)

	err = app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}

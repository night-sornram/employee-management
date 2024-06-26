package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/employee-service/adapter"
	"github.com/night-sornram/employee-management/employee-service/middleware"
	"github.com/night-sornram/employee-management/employee-service/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	const (
		host     = "db"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "employee"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// amqp := adapter.MessagingClient{}
	// amqp.ConnectToBroker("amqp://user:password@rabbitmq/")
	// defer amqp.Close()

	repo := adapter.NewGormAdapter(db)
	service := repository.NewEmployeeService(repo)
	handle := adapter.NewHandlerFiber(service)
	// handle := adapter.NewHandleFiber(service, &amqp)

	err = db.AutoMigrate(&repository.Employee{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	app.Use(cors.New())

	app.Post("/login", handle.Login)
	app.Post("/logout", handle.Logout)
	app.Get("/me", handle.GetMe)

	app.Use("/api", middleware.Protected())
	app.Get("/api/employees", middleware.Authorize("admin"), handle.GetEmployees)
	app.Get("/api/employees/download", middleware.Authorize("admin"), handle.DownloadCSV)
	app.Get("/api/employees/:id", middleware.Authorize("admin", "user"), handle.GetEmployee)
	app.Post("/api/employees", middleware.Authorize("admin", "user"), handle.CreateEmployee)
	app.Put("/api/employees/:id", middleware.Authorize("admin", "user"), handle.UpdateEmployee)
	app.Delete("/api/employees/:id", middleware.Authorize("admin"), handle.DeleteEmployee)
	app.Post("/api/changePassword", handle.ChangePassword)

	err = app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/adapter"
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

	app.Get("/employees", handle.GetEmployees)
	app.Get("/employees/:id", handle.GetEmployee)
	app.Post("/employees", handle.CreateEmployee)
	app.Put("/employees/:id", handle.UpdateEmployee)
	app.Delete("/employees/:id", handle.DeleteEmployee)

	app.Listen(":8080")
}

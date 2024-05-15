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
		dbname   = "postgres"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Leave{})
	repo := adapter.NewGormAdapter(db)
	service := repository.NewLeaveService(repo)
	handle := adapter.NewhandlerFiber(service)

	app.Get("/leaves", handle.GetLeaves)
	app.Get("/leaves/:id", handle.GetLeave)
	app.Post("/leaves", handle.CreateLeave)
	app.Put("/leaves/:id", handle.UpdateLeave)
	app.Delete("/leaves/:id", handle.DeleteLeave)
	app.Put("/leaves/approval/:id", handle.UpdateStatus)

	app.Listen(":8082")
}

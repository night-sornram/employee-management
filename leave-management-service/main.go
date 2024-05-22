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
		dbname   = "leave"
		tz       = "Asia/Bangkok"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, user, password, dbname, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Leave{})
	repo := adapter.NewGormAdapter(db)
	service := repository.NewLeaveService(repo)
	handle := adapter.NewhandlerFiber(service)

	app.Use(cors.New())

	app.Use("/api", middleware.Protected())

	app.Get("/api/leaves", handle.GetLeaves)
	app.Get("/api/leaves/me/:eid", handle.GetMyLeaves)
	app.Get("/api/leaves/:id", handle.GetLeave)
	app.Post("/api/leaves", handle.CreateLeave)
	app.Put("/api/leaves/:id", handle.UpdateLeave)
	app.Delete("/api/leaves/:id", handle.DeleteLeave)

	app.Listen(":8082")
}

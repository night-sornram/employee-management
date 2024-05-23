package main

import (
	"fmt"
	"log"

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
		dbname   = "postgres"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&repository.Attendance{})

	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	repo := adapter.NewGormAdapter(db)
	service := repository.NewAttendanceService(repo)
	handle := adapter.NewhandlerFiber(service)

	app.Use(cors.New())
	app.Use("/api", middleware.Protected())

	app.Get("/api/attendances", handle.GetAttendances)
	app.Get("/api/attendances/:id", handle.GetAttendance)
	app.Post("/api/attendances", handle.CreateAttendance)
	app.Post("/api/attendances/check-in", handle.CheckIn)
	app.Put("/api/attendances/check-out", handle.CheckOut)
	app.Put("/api/attendances/:id", handle.UpdateAttendance)
	app.Delete("/api/attendances/:id", handle.DeleteAttendance)
	app.Get("/api/attendances/me/:eid", handle.GetMyAttendances)
	app.Get("/api/attendances/check-today/:eid", handle.CheckToday)

	app.Listen(":8081")
}

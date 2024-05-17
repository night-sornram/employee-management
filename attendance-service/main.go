package main

import (
	"fmt"
	"log"

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

	err = db.AutoMigrate(&repository.Attendance{})
	
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	repo := adapter.NewGormAdapter(db)
	service := repository.NewAttendanceService(repo)
	handle := adapter.NewhandlerFiber(service)

	app.Get("/attendance", handle.GetAttendances)
	app.Get("/attendance/:id", handle.GetAttendance)
	app.Post("/attendance", handle.CreateAttendance)
	app.Post("/attendance/check-in", handle.CheckIn)
	app.Put("/attendance/check-out", handle.CheckOut)
	app.Put("/attendance/:id", handle.UpdateAttendance)
	app.Delete("/attendance/:id", handle.DeleteAttendance)

	app.Listen(":8081")
}

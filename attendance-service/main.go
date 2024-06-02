package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/adapter"
	"github.com/night-sornram/employee-management/common_utils"
	"github.com/night-sornram/employee-management/common_utils/middleware"
	"github.com/night-sornram/employee-management/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()

	dsn := common_utils.ConnectDB("8081")
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
	handle := adapter.NewHandlerFiber(service)

	app.Use(cors.New())
	app.Use("/api/attendances", middleware.Protected())

	app.Get("/api/attendances", handle.GetAttendances)
	app.Get("/api/attendances/:id", handle.GetAttendance)
	app.Post("/api/attendances", handle.CreateAttendance)
	app.Post("/api/attendances/check-in", handle.CheckIn)
	app.Put("/api/attendances/check-out", handle.CheckOut)
	app.Put("/api/attendances/:id", handle.UpdateAttendance)
	app.Delete("/api/attendances/:id", handle.DeleteAttendance)
	app.Get("/api/attendances/me/:eid", handle.GetMyAttendances)
	app.Get("/api/attendances/check-today/:eid", handle.CheckToday)

	err = app.Listen(":8081")
	if err != nil {
		return
	}
}

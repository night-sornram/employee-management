package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/attendance-service/adapter"
	"github.com/night-sornram/employee-management/attendance-service/middleware"
	"github.com/night-sornram/employee-management/attendance-service/repository"
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
		dbname   = "attendance"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}

	err = db.AutoMigrate(&repository.Attendance{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	extension := "dblink"
	createExtension := fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %s;", extension)
	db.Exec(createExtension)

	repo := adapter.NewGormAdapter(db)
	service := repository.NewAttendanceService(repo)
	handle := adapter.NewHandlerFiber(service)

	app.Use(cors.New())
	app.Use("/api", middleware.Protected())

	app.Get("/api/attendances", handle.GetAttendances)
	app.Get("/api/attendances/me/:eid", handle.GetMyAttendances)
	app.Get("/api/attendances/check-today/:eid", handle.CheckToday)
	app.Get("/api/attendances/download", middleware.Authorize("admin"), handle.DownloadCSV)
	app.Post("/api/attendances", handle.CreateAttendance)
	app.Post("/api/attendances/check-in", handle.CheckIn)
	app.Put("/api/attendances/check-out", handle.CheckOut)
	app.Put("/api/attendances/:id", handle.UpdateAttendance)
	app.Delete("/api/attendances/:id", handle.DeleteAttendance)
	app.Get("/api/attendances/late/day", handle.GetDayLate)
	app.Get("/api/attendances/late/month", handle.GetMonthLate)
	app.Get("/api/attendances/late/year/:year", handle.GetYearLate)
	app.Get("/api/attendances/late/all", handle.GetAllLate)
	app.Get("/api/attendances/:id", handle.GetAttendance)

	err = app.Listen("0.0.0.0:8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

}

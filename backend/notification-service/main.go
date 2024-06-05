package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/notification-service/adapter"
	"github.com/night-sornram/employee-management/notification-service/middleware"
	"github.com/night-sornram/employee-management/notification-service/repository"
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
		dbname   = "notification"
		tz       = "Asia/Bangkok"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, user, password, dbname, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&repository.Notification{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	repo := adapter.NewGormAdapter(db)
	service := repository.NewNotificationService(repo)
	handle := adapter.NewhandlerFiber(service)

	app.Use(cors.New())

	app.Use("/api", middleware.Protected())

	app.Get("/api/notifications", handle.GetNotifications)
	app.Get("/api/notifications/employee/:employeeID", handle.GetNotificationByEmployeeID)
	app.Get("/api/notifications/:id", handle.GetNotification)
	app.Post("/api/notifications", handle.CreateNotification)
	app.Put("/api/notifications/read/:id", handle.ReadNotification)
	app.Put("/api/notifications/:id", handle.UpdateNotification)
	app.Delete("/api/notifications/:id", handle.DeleteNotification)
	err = app.Listen("0.0.0.0:8083")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/common-utils"
	"github.com/night-sornram/employee-management/common-utils/middleware"
	"github.com/night-sornram/employee-management/notification-service/adapter"
	"github.com/night-sornram/employee-management/notification-service/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()

	dsn := common_utils.ConnectDB("8083")
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

	app.Use("/api/notifications", middleware.Protected())

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

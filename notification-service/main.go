package main

import (
	"employee/adapter"
	"employee/repository"
	"github.com/night-sornram/employee-management/common_utils"
	"github.com/night-sornram/employee-management/common_utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := common_utils.ConnectDB("8083")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Notification{})
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

	app.Listen(":8083")
}

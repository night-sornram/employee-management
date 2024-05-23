package main

import (
	"fmt"

	"employee/adapter"
	"employee/middleware"
	"employee/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		dbname   = "notification"
		tz       = "Asia/Bangkok"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, user, password, dbname, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Notification{})
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

	app.Listen(":8083")
}

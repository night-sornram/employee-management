package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/leave-management-service/adapter"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"github.com/webbalaka/common-utils"
	"github.com/webbalaka/common-utils/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()

	dsn := common_utils.ConnectDB("8082")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&repository.Leave{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	extension := "dblink"
	createExtension := fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %s;", extension)
	db.Exec(createExtension)

	repo := adapter.NewGormAdapter(db)
	service := repository.NewLeaveService(repo)
	handle := adapter.NewHandlerFiber(service)

	app.Use(cors.New())

	app.Use("/api/leaves", middleware.Protected())

	app.Get("/api/leaves", handle.GetLeaves)
	app.Get("/api/leaves/me/:eid", handle.GetAllMe)
	app.Get("/api/leaves/:id", handle.GetLeave)
	app.Post("/api/leaves", handle.CreateLeave)
	app.Put("/api/leaves/:id", handle.UpdateLeave)
	app.Delete("/api/leaves/:id", handle.DeleteLeave)
	app.Put("/api/leaves/approval/:id", handle.UpdateStatus)

	err = app.Listen("0.0.0.0:8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

}

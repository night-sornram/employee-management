package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/night-sornram/employee-management/adapter"
	"github.com/night-sornram/employee-management/common_utils"
	"github.com/night-sornram/employee-management/common_utils/middleware"
	"github.com/night-sornram/employee-management/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setup() *fiber.App {
	app := fiber.New()

	dsn := common_utils.ConnectDB("8082")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Leave{})

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

	app.Listen(":8082")

	return app
}

func main() {
	app := setup()
	app.Listen(":8082")
}

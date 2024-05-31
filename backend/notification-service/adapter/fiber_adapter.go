package adapter

import (
	"github.com/night-sornram/employee-management/notification-service/repository"

	"github.com/gofiber/fiber/v2"
)

type HandlerFiber struct {
	service repository.NotificationService
}

func NewhandlerFiber(service repository.NotificationService) HandlerFiber {
	return HandlerFiber{
		service: service,
	}
}

func (f *HandlerFiber) GetNotifications(c *fiber.Ctx) error {
	notifications, err := f.service.GetNotifications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(notifications)
}

func (f *HandlerFiber) GetNotification(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	notification, err := f.service.GetNotification(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(notification)
}

func (f *HandlerFiber) CreateNotification(c *fiber.Ctx) error {
	var notification repository.Notification
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newNotification, err := f.service.CreateNotification(notification)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(newNotification)
}

func (f *HandlerFiber) UpdateNotification(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var notification repository.Notification
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	notification, err = f.service.UpdateNotification(id, notification)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(notification)
}

func (f *HandlerFiber) DeleteNotification(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = f.service.DeleteNotification(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Notification deleted successfully",
	})
}

func (f *HandlerFiber) GetNotificationByEmployeeID(c *fiber.Ctx) error {
	employeeID := c.Params("employeeID")
	notifications, err := f.service.GetNotificationByEmployeeID(employeeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(notifications)
}

func (f *HandlerFiber) ReadNotification(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = f.service.ReadNotification(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Notification read successfully",
	})
}

package adapter

import (
	"employee/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handlerFiber struct {
	service repository.NotificationService
}

func NewhandlerFiber(service repository.NotificationService) handlerFiber {
	return handlerFiber{
		service: service,
	}
}

func (f *handlerFiber) GetNotifications(c *fiber.Ctx) error {
	notifications, err := f.service.GetNotifications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(notifications)
}

func (f *handlerFiber) GetNotification(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusOK).JSON(notification)
}

func (f *handlerFiber) CreateNotification(c *fiber.Ctx) error {
	var notification repository.Notification
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(notification)
	if err != nil {
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
	return c.Status(fiber.StatusCreated).JSON(newNotification)
}

func (f *handlerFiber) UpdateNotification(c *fiber.Ctx) error {
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

	validate := validator.New()
	err = validate.Struct(notification)
	if err != nil {
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

func (f *handlerFiber) DeleteNotification(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Notification deleted successfully",
	})
}

func (f *handlerFiber) GetNotificationByEmployeeID(c *fiber.Ctx) error {
	employeeID := c.Params("employeeID")
	notifications, err := f.service.GetNotificationByEmployeeID(employeeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(notifications)
}

func (f *handlerFiber) ReadNotification(c *fiber.Ctx) error {
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

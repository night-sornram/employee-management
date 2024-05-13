package adapter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/night-sornram/employee-management/repository"
)

type handlerFiber struct {
	service repository.AttendanceService
}

func NewhandlerFiber(service repository.AttendanceService) handlerFiber {
	return handlerFiber{
		service: service,
	}
}

func (f *handlerFiber) GetAttendances(c *fiber.Ctx) error {
	attendances, err := f.service.GetAttendances()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(attendances)
}

func (f *handlerFiber) GetAttendance(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	attendance, err := f.service.GetAttendance(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(attendance)
}

func (f *handlerFiber) CreateAttendance(c *fiber.Ctx) error {
	var attendance repository.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newAttendance, err := f.service.CreateAttendance(attendance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(newAttendance)
}

func (f *handlerFiber) UpdateAttendance(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var attendance repository.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	updatedAttendance, err := f.service.UpdateAttendance(id, attendance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(updatedAttendance)
}

func (f *handlerFiber) DeleteAttendance(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = f.service.DeleteAttendance(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

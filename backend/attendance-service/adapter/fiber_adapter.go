package adapter

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/attendance-service/repository"
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
	return c.Status(fiber.StatusOK).JSON(attendances)
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
	return c.Status(fiber.StatusOK).JSON(attendance)
}

func (f *handlerFiber) CreateAttendance(c *fiber.Ctx) error {
	var attendance repository.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(attendance)
	if err != nil {
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
	return c.Status(fiber.StatusCreated).JSON(newAttendance)
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

	validate := validator.New()
	err = validate.Struct(attendance)
	if err != nil {
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
	return c.Status(fiber.StatusOK).JSON(updatedAttendance)
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

func (f *handlerFiber) CheckIn(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newCheckIn, err := f.service.CheckIn(data["eid"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newCheckIn)
}

func (f *handlerFiber) CheckOut(c *fiber.Ctx) error {
	var data map[string]int
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newCheckOut, err := f.service.CheckOut(data["id"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(newCheckOut)
}

func (f *handlerFiber) GetMyAttendances(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	attendances, err := f.service.GetMyAttendances(eid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (f *handlerFiber) CheckToday(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	attendance, _ := f.service.CheckToday(eid)
	if attendance.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(nil)
	}
	return c.Status(fiber.StatusOK).JSON(attendance)
}

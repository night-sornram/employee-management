package adapter

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/repository"
)

type handlerFiber struct {
	service repository.LeaveService
}

func NewhandlerFiber(service repository.LeaveService) handlerFiber {
	return handlerFiber{
		service: service,
	}
}

func (f *handlerFiber) GetLeaves(c *fiber.Ctx) error {
	leaves, err := f.service.GetLeaves()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(leaves)
}

func (f *handlerFiber) GetLeave(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	leave, err := f.service.GetLeave(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(leave)
}

func (f *handlerFiber) CreateLeave(c *fiber.Ctx) error {
	var leave repository.Leave
	if err := c.BodyParser(&leave); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(leave)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newLeave, err := f.service.CreateLeave(leave)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newLeave)
}

func (f *handlerFiber) UpdateLeave(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var leave repository.Leave
	if err := c.BodyParser(&leave); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validate := validator.New()
	// err = validate.Struct(leave)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }
	updateLeave, err := f.service.UpdateLeave(id, leave)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updateLeave)
}

func (f *handlerFiber) DeleteLeave(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = f.service.DeleteLeave(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (f *handlerFiber) UpdateStatus(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var leave repository.LeaveStatus
	if err := c.BodyParser(&leave); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err = validate.Struct(leave)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	updateStatus, err := f.service.UpdateStatus(id, leave)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updateStatus)
}

func (f *handlerFiber) GetMyLeaves(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	leaves, err := f.service.GetMyLeaves(eid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(leaves)
}

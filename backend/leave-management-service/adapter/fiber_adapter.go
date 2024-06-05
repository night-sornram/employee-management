package adapter

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"strconv"
)

type handlerFiber struct {
	service repository.LeaveService
}

func NewHandlerFiber(service repository.LeaveService) handlerFiber {
	return handlerFiber{
		service: service,
	}
}

func (f *handlerFiber) GetLeaves(c *fiber.Ctx) error {
	query := repository.Query{
		Date:    "",
		Page:    1,
		Name:    "",
		PerPage: 8,
		Status:  "",
		Option:  "",
	}

	if d := c.Query("date"); d != "" {
		query.Date = d
	}

	if n := c.Query("name"); n != "" {
		query.Name = n
	}

	if s := c.Query("status"); s != "" {
		query.Status = s
	}

	if o := c.Query("option"); o != "" {
		query.Option = o
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	query.Page = page

	leaves, err := f.service.GetLeaves(query)
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

func (f *handlerFiber) GetAllMe(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	query := repository.Query{
		Date:    "",
		Page:    1,
		Name:    "",
		PerPage: 8,
		Status:  "",
		Option:  "",
	}

	if d := c.Query("date"); d != "" {
		query.Date = d
	}

	if o := c.Query("option"); o != "" {
		query.Option = o
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	query.Page = page

	leaves, err := f.service.GetAllMe(query, eid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(leaves)
}

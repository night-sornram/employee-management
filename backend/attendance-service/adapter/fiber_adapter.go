package adapter

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/attendance-service/repository"
)

type handlerFiber struct {
	service repository.AttendanceService
}

func NewHandlerFiber(service repository.AttendanceService) handlerFiber {
	return handlerFiber{
		service: service,
	}
}

func (f *handlerFiber) GetAttendances(c *fiber.Ctx) error {
	query := repository.Query{
		Date:    "",
		Page:    1,
		Name:    "",
		PerPage: 8,
		Option:  "",
		LeaveID: 0,
	}

	if d := c.Query("date"); d != "" {
		query.Date = d
	}

	if n := c.Query("name"); n != "" {
		query.Name = n
	}

	if o := c.Query("option"); o != "" {
		query.Option = o
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	query.Page = page

	leaveID, _ := strconv.Atoi(c.Query("leave_id", "0"))
	query.LeaveID = leaveID

	attendances, err := f.service.GetAttendances(query)

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

	query := repository.Query{
		Date:    "",
		Page:    1,
		Name:    "",
		PerPage: 8,
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

	data, err := f.service.GetMyAttendances(query, eid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func (f *handlerFiber) CheckToday(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	attendance, err := f.service.CheckToday(eid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if attendance.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(nil)
	}
	return c.Status(fiber.StatusOK).JSON(attendance)
}

func (f *handlerFiber) GetDayLate(c *fiber.Ctx) error {
	attendances, err := f.service.GetDayLate()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (f *handlerFiber) GetMonthLate(c *fiber.Ctx) error {
	var date repository.GetMonth
	if err := c.BodyParser(&date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	attendances, err := f.service.GetMonthLate(date)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (f *handlerFiber) GetYearLate(c *fiber.Ctx) error {
	year, err := c.ParamsInt("year")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	attendances, err := f.service.GetYearLate(year)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (f *handlerFiber) GetAllLate(c *fiber.Ctx) error {
	attendances, err := f.service.GetAllLate()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (f *handlerFiber) DownloadCSV(c *fiber.Ctx) error {
	query := c.Query("query")
	//if query == "" {
	//	return c.Status(fiber.StatusBadRequest).SendString("Query is missing")
	//}
	data, err := f.service.DownloadCSV(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Nope")
	}
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=data.csv")
	c.Set(fiber.HeaderContentType, "text/csv")
	return c.Send(data)
}

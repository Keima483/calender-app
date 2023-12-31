package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/models"
	"github.com/keima483/calender-app/services"
)

func GetHolidays(c *fiber.Ctx) error {
	return c.JSON(services.GetHolidays())
}

func GetTasks(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Locals("user_id").(string))
	tasks, err := services.GetTasks(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(tasks)
}

func AddHolidays(c *fiber.Ctx) error {
	var holidays []*models.HolidayModel
	if err := c.BodyParser(&holidays); err != nil {
		return err
	}
	hol := services.AddHolidays(holidays)
	return c.JSON(hol)
}

func AddTask(c *fiber.Ctx) error {
	var task models.TaskModel
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Locals("user_id").(string))
	tasks, err := services.AddTask(task, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(tasks)
}

func AddMultipleTasks(c *fiber.Ctx) error {
	var tasks []models.TaskModel
	if err := c.BodyParser(&tasks); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Locals("user_id").(string))
	tsks, err := services.AddTasks(tasks, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(tsks)
}

func GetDayDetail(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Locals("user_id").(string))
	date := c.Query("date")
	task, holidays, err := services.GetDayDetail(date, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(fiber.Map{"tasks":task, "holidays": holidays})
}

func DeleteTask(c *fiber.Ctx) error {
	taskId, _ := strconv.Atoi(c.Params("id"))
	err := services.DeleteTask(taskId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.SendString("Successfully deleted")
}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/controllers"
	"github.com/keima483/calender-app/middleware"
)

func APIRoutes(app *fiber.App) {
	group := app.Group("/api/v1/gen")
	group.Use(middleware.VerifyJWTToken)

	group.Post("/holiday", controllers.AddHolidays)
	group.Get("/holiday", controllers.GetHolidays)
	group.Get("/task", controllers.GetTasks)
	group.Post("/task", controllers.AddTask)
	group.Post("/tasks", controllers.AddMultipleTasks)
	group.Delete("/task/:id", controllers.DeleteTask)
}
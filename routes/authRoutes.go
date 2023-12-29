package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/controllers"
)

func AuthRoutes(app *fiber.App) {
	group := app.Group("/api/v1/user")
	group.Post("/login", controllers.Login)
	group.Post("/signup", controllers.SignUp)
}

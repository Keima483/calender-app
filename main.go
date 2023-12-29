package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/initializers"
	"github.com/keima483/calender-app/routes"
)

func init() {
	config:= initializers.LoadConfig(".")
	log.Println(config)
	if err := initializers.InitialiseDB(&config); err != nil {
		panic(err.Error())
	}
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Calender app service started running")
	})
	routes.AuthRoutes(app)
	routes.APIRoutes(app)
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}

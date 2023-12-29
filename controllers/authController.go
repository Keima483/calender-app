package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/models"
	"github.com/keima483/calender-app/services"
	"github.com/keima483/calender-app/middleware"
)

func Login(c *fiber.Ctx) error {
	loginDetails := new(models.LoginModel)
	if err := c.BodyParser(loginDetails); err != nil {
		return err
	}
	user, err := services.Login(loginDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	token, exp, err := middleware.CreateJWTToken(user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func SignUp(c *fiber.Ctx) error {
	userDetails := new(models.UserModel)
	if err := c.BodyParser(userDetails); err != nil {
		return err
	}
	user, err := services.SignUp(userDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} 
	token, exp, err := middleware.CreateJWTToken(user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}


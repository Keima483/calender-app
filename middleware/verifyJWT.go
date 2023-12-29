package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keima483/calender-app/initializers"
)

func VerifyJWTToken(c *fiber.Ctx) error {

	var tokenString string
	authorization := c.Get("Authorization")
	if !(strings.HasPrefix(authorization, "Bearer")) {
		return c.Status(400).JSON(fiber.Map{"message": "JWT token Not provided or is not in correct systax"})
	} 
	tokenString = strings.Split(authorization, " ")[1]
	config := initializers.LoadConfig(".")
	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid JWT token"})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid claims"})
	}
	c.Locals("user_id", fmt.Sprint(claims["user_id"]))
	return c.Next()
}
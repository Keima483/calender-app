package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/keima483/calender-app/initializers"
	"github.com/keima483/calender-app/repository"
)

func CreateJWTToken(user repository.User) (string, int64, error) {
	config := initializers.LoadConfig(".")
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

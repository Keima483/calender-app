package services

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/calender-app/models"
	"github.com/keima483/calender-app/initializers"
	"github.com/keima483/calender-app/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(lm *models.LoginModel) (repository.User, error) {
	user := new(repository.User)
	initializers.DB.Where("email = ?", lm.Email).First(&user)
	if user.Email == "" {
		return repository.User{}, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lm.Password)); err != nil {
		return repository.User{}, fiber.NewError(fiber.StatusBadRequest, "Incorrect Password")
	}
	return *user, nil
}

func SignUp(um *models.UserModel) (repository.User, error) {
	if um.Email == "" && um.Password == "" {
		return repository.User{}, errors.New("enter email and password atleast")
	}
	user := repository.UserFromModel(*um)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return repository.User{}, errors.New("not able to encrypt password")
	}
	user.Password = string(hash)
	initializers.DB.Save(&user)
	return user, nil
}

package controller

import (
	"moview/src/models"
	"moview/src/repository"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthRepository repository.AuthRepository
}

func NewAuthController(authRepository repository.AuthRepository) *AuthController {
	return &AuthController{AuthRepository: authRepository}
}

func (ac *AuthController) LogIn(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	u, err := ac.AuthRepository.LogIn(user)
	if err != nil {
		return err
	}
	return c.JSON(u)
}
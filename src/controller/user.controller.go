package controller

import (
	"moview/src/models"
	"moview/src/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{UserRepository: userRepository}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if err := uc.UserRepository.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(user)
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user, err := uc.UserRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.ID = id

	if err := uc.UserRepository.UpdateUser(user); err != nil {
		return err
	}

	return c.JSON(user)
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err := uc.UserRepository.DeleteUser(id); err != nil {
		return err
	}

	return c.SendString("User deleted successfully")
}
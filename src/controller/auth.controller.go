package controller

import (
	"moview/src/repository"
)

type AuthController struct {
	AuthRepository repository.AuthRepository
}

func NewAuthController(authRepository repository.AuthRepository) *AuthController {
	return &AuthController{AuthRepository: authRepository}
}

package user_controller

import "github.com/tiagocosta/auction-app/internal/usecase/user_usecase"

type UserController struct {
	UserUseCase user_usecase.UserUseCase
}

func NewUserController(userUseCase user_usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

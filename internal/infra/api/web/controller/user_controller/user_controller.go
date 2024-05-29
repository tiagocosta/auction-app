package user_controller

import "github.com/tiagocosta/auction-app/internal/usecase/user_usecase"

type UserController struct {
	UserUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

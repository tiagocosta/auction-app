package user_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (uc *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	user, err := uc.UserRepositry.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

package user_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/entity/user_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type UserUseCase struct {
	UserRepositry user_entity.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

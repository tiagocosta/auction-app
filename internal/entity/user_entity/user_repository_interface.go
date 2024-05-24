package user_entity

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}

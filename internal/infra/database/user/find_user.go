package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/user_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}
	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			message := fmt.Sprintf("user not found with id %s", userId)
			logger.Error(message, err)
			return nil, internal_error.NewNotFoundError(message)
		}

		message := "error trying to find user by userid"
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}

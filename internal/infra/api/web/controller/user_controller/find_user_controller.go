package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiagocosta/auction-app/configuration/rest_err"
)

func (controller *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields", rest_err.Cause{
			Field:   "userId",
			Message: "invalid uui value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := controller.UserUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.FromInternalError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}

package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiagocosta/auction-app/configuration/rest_err"
	"github.com/tiagocosta/auction-app/internal/infra/api/web/validation"
	"github.com/tiagocosta/auction-app/internal/usecase/bid_usecase"
)

func (controller *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO
	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := controller.BidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := rest_err.FromInternalError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}

package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiagocosta/auction-app/configuration/rest_err"
)

func (controller *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields", rest_err.Cause{
			Field:   "auctionId",
			Message: "invalid uui value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	bidsData, err := controller.BidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.FromInternalError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidsData)
}

package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiagocosta/auction-app/configuration/rest_err"
	"github.com/tiagocosta/auction-app/internal/usecase/auction_usecase"
)

func (controller *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields", rest_err.Cause{
			Field:   "auctionId",
			Message: "invalid uui value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := controller.AuctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.FromInternalError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (controller *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("error trying to validate auction status param")
		c.JSON(errRest.Code, errRest)
		return
	}
	statusValue := auction_usecase.AuctionStatus(statusNumber)

	auctions, err := controller.AuctionUseCase.FindAuctions(context.Background(), statusValue, category, productName)
	if err != nil {
		errRest := rest_err.FromInternalError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (controller *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields", rest_err.Cause{
			Field:   "auctionId",
			Message: "invalid uui value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := controller.AuctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.FromInternalError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

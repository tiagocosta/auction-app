package auction_controller

import (
	"github.com/tiagocosta/auction-app/internal/usecase/auction_usecase"
)

type AuctionController struct {
	AuctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *AuctionController {
	return &AuctionController{
		AuctionUseCase: auctionUseCase,
	}
}

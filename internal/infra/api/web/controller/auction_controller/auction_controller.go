package auction_controller

import (
	"github.com/tiagocosta/auction-app/internal/usecase/auction_usecase"
)

type AuctionController struct {
	AuctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		AuctionUseCase: auctionUseCase,
	}
}

package bid_controller

import "github.com/tiagocosta/auction-app/internal/usecase/bid_usecase"

type BidController struct {
	BidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		BidUseCase: bidUseCase,
	}
}

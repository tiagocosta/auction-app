package bid_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (uc *BidUseCase) CreateBid(ctx context.Context, bidInput BidInputDTO) *internal_error.InternalError {
	bid, err := bid_entity.NewBid(bidInput.UserId, bidInput.AuctionId, bidInput.Amount)
	if err != nil {
		return err
	}

	uc.BidChannel <- *bid

	return nil
}

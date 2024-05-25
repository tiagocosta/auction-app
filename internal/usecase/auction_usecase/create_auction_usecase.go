package auction_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (uc *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionIntputDTO) *internal_error.InternalError {
	auction, err := auction_entity.NewAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	err = uc.AuctionRepository.CreateAuction(ctx, *auction)
	if err != nil {
		return err
	}

	return nil
}

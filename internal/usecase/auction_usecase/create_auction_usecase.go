package auction_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/configuration/logger"
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

	err = uc.AuctionRepository.CreateAuction(ctx, auction)
	if err != nil {
		return err
	}

	uc.startAuctionRoutine(ctx, auction)

	return nil
}

func (uc *AuctionUseCase) startAuctionRoutine(ctx context.Context, auction *auction_entity.Auction) {
	go func() {
		for {
			select {
			case <-uc.Timer.C:
				if auction.IsActive() {
					auction.Finish()
					err := uc.AuctionRepository.UpdateAuctionStatus(ctx, auction)
					if err != nil {
						logger.Error("error trying to update auction status", err)
					}
					uc.Timer.Reset(uc.Interval)
				}
			}
		}
	}()
}

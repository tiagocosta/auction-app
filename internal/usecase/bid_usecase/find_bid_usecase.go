package bid_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (uc *BidUseCase) FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bids, err := uc.BidRepository.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	var bidsOutput []BidOutputDTO
	for _, bid := range bids {
		bidsOutput = append(bidsOutput, BidOutputDTO{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidsOutput, nil
}

func (uc *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bid, err := uc.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	bidOutput := &BidOutputDTO{
		Id:        bid.Id,
		UserId:    bid.UserId,
		AuctionId: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}

	return bidOutput, nil
}

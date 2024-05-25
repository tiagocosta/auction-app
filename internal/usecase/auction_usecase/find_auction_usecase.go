package auction_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (uc *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auction, err := uc.AuctionRepository.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}, nil
}

func (uc *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {
	auctions, err := uc.AuctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionsOutput []AuctionOutputDTO
	for _, auction := range auctions {
		auctionsOutput = append(auctionsOutput, AuctionOutputDTO{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}

	return auctionsOutput, nil
}

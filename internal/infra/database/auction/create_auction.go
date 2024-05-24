package auction

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		return internal_error.NewInternalServerError("error trying to insert auction")
	}

	return nil
}

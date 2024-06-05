package auction_entity

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]Auction, *internal_error.InternalError)
	UpdateAuctionStatus(ctx context.Context, auction *Auction) *internal_error.InternalError
}

package bid_entity

import (
	"context"

	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bids []Bid) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}

package bid_usecase

import (
	"context"
	"time"

	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidRepositoryInterface
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bids []bid_entity.Bid) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
}

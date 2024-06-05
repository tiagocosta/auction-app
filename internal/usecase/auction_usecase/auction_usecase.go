package auction_usecase

import (
	"context"
	"os"
	"time"

	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
	"github.com/tiagocosta/auction-app/internal/usecase/bid_usecase"
)

type ProductCondition int64

type AuctionStatus int64

type AuctionIntputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
	BidRepository     bid_entity.BidRepositoryInterface

	Timer    *time.Timer
	Interval time.Duration
}

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInput AuctionIntputDTO) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

func NewAuctionUseCase(auctionRepository auction_entity.AuctionRepositoryInterface, bidRepository bid_entity.BidRepositoryInterface) AuctionUseCaseInterface {
	auctionInterval := getAuctionInterval()

	return &AuctionUseCase{
		AuctionRepository: auctionRepository,
		BidRepository:     bidRepository,
		Timer:             time.NewTimer(auctionInterval),
		Interval:          auctionInterval,
	}
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return 60 * time.Second
	}
	return duration
}

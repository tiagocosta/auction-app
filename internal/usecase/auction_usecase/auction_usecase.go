package auction_usecase

import (
	"context"
	"time"

	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type ProductCondition int64

type AuctionStatus int64

type AuctionIntputDTO struct {
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
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

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
}

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInput AuctionIntputDTO) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
}

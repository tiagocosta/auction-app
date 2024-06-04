package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type BidInputDTO struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidRepositoryInterface

	Timer               *time.Timer
	MaxBatchSize        int
	BatchInsertInterval time.Duration
	BidChannel          chan bid_entity.Bid
}

func NewBidUseCase(bidRepository bid_entity.BidRepositoryInterface) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	uc := &BidUseCase{
		BidRepository:       bidRepository,
		Timer:               time.NewTimer(maxSizeInterval),
		MaxBatchSize:        maxBatchSize,
		BatchInsertInterval: maxSizeInterval,
		BidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	uc.startBidBatchRoutine(context.Background())

	return uc
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration
}

func getMaxBatchSize() int {
	maxBatchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}
	return maxBatchSize
}

var bidBatch []bid_entity.Bid

func (uc *BidUseCase) startBidBatchRoutine(ctx context.Context) {
	go func() {
		defer close(uc.BidChannel)
		for {
			select {
			case bid, ok := <-uc.BidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						err := uc.BidRepository.CreateBid(ctx, bidBatch)
						if err != nil {
							logger.Error("error trying to process bid batch list", err)
						}
					}
					return
				}
				bidBatch = append(bidBatch, bid)
				if len(bidBatch) > uc.MaxBatchSize {
					err := uc.BidRepository.CreateBid(ctx, bidBatch)
					if err != nil {
						logger.Error("error trying to process bid batch list", err)
					}
					bidBatch = nil
					uc.Timer.Reset(uc.BatchInsertInterval)
				}
			case <-uc.Timer.C:
				err := uc.BidRepository.CreateBid(ctx, bidBatch)
				if err != nil {
					logger.Error("error trying to process bid batch list", err)
				}
				bidBatch = nil
				uc.Timer.Reset(uc.BatchInsertInterval)
			}
		}
	}()
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInput BidInputDTO) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
}

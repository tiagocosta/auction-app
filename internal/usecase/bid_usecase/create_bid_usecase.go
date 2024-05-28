package bid_usecase

import (
	"context"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

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

func (uc *BidUseCase) CreateBid(ctx context.Context, bidInput BidInputDTO) *internal_error.InternalError {
	bid, err := bid_entity.NewBid(bidInput.UserId, bidInput.AuctionId, bidInput.Amount)
	if err != nil {
		return err
	}

	uc.BidChannel <- *bid

	return nil
}

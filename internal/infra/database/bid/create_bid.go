package bid

import (
	"context"
	"sync"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

func (br *BidRepository) CreateBid(ctx context.Context, bids []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bids {
		wg.Add(1)
		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auction, err := br.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)
			if err != nil {
				message := "error trying to find auction by id"
				logger.Error(message, err)
				return
			}

			if !auction.IsActive() {
				return
			}

			bidMongo := &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			_, errMongo := br.Collection.InsertOne(ctx, bidMongo)
			if errMongo != nil {
				message := "error trying to insert bid"
				logger.Error(message, err)
				return
			}
		}(bid)
	}

	wg.Wait()
	return nil
}

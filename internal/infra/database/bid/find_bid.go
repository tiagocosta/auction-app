package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/bid_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}

	cursor, err := br.Collection.Find(ctx, filter)
	if err != nil {
		message := fmt.Sprintf("error trying to find bids by auction id %s", auctionId)
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	var bidsMongo []BidEntityMongo
	err = cursor.All(ctx, &bidsMongo)
	if err != nil {
		message := fmt.Sprintf("error trying to find bids by auction id %s", auctionId)
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	var bids []bid_entity.Bid
	for _, bidMongo := range bidsMongo {
		bids = append(bids, bid_entity.Bid{
			Id:        bidMongo.Id,
			UserId:    bidMongo.UserId,
			AuctionId: bidMongo.AuctionId,
			Amount:    bidMongo.Amount,
			Timestamp: time.Unix(bidMongo.Timestamp, 0),
		})
	}

	return bids, nil
}

func (br *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auctionId": auctionId}
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	var bidMongo BidEntityMongo
	err := br.Collection.FindOne(ctx, filter, opts).Decode(&bidMongo)
	if err != nil {
		message := fmt.Sprintf("error trying to find winner bid by auction id %s", auctionId)
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	return &bid_entity.Bid{
		Id:        bidMongo.Id,
		UserId:    bidMongo.UserId,
		AuctionId: bidMongo.AuctionId,
		Amount:    bidMongo.Amount,
		Timestamp: time.Unix(bidMongo.Timestamp, 0),
	}, nil
}

package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"github.com/tiagocosta/auction-app/internal/entity/auction_entity"
	"github.com/tiagocosta/auction-app/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionMongo AuctionEntityMongo
	err := ar.Collection.FindOne(ctx, filter).Decode(&auctionMongo)
	if err != nil {
		message := fmt.Sprintf("error trying to find auction by id %s", id)
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	return &auction_entity.Auction{
		Id:          auctionMongo.Id,
		ProductName: auctionMongo.ProductName,
		Category:    auctionMongo.Category,
		Description: auctionMongo.Description,
		Condition:   auctionMongo.Condition,
		Status:      auctionMongo.Status,
		Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category string,
	productName string) ([]auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		message := "error trying to find auctions"
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}
	defer cursor.Close(ctx)

	var auctionsMongo []AuctionEntityMongo
	err = cursor.All(ctx, auctionsMongo)
	if err != nil {
		message := "error trying to find auctions"
		logger.Error(message, err)
		return nil, internal_error.NewInternalServerError(message)
	}

	var auctions []auction_entity.Auction
	for _, auctionMongo := range auctionsMongo {
		auctions = append(auctions, auction_entity.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctions, nil
}

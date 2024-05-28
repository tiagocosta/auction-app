package bid_entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

func NewBid(userId, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.NewString(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	err := bid.Validate()
	if err != nil {
		return nil, err
	}
	return bid, nil
}

func (bid *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(bid.Id); err != nil {
		return internal_error.NewBadRequestError("invalid user id")
	}
	if err := uuid.Validate(bid.AuctionId); err != nil {
		return internal_error.NewBadRequestError("invalid auction id")
	}
	if bid.Amount <= 0 {
		return internal_error.NewBadRequestError("invalid amount value")
	}
	return nil
}

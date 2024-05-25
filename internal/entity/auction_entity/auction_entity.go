package auction_entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/tiagocosta/auction-app/internal/internal_error"
)

type ProductCondition int

const (
	New ProductCondition = iota
	Used
	Refurbished
)

type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

func NewAuction(productName, category, description string, condition ProductCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.NewString(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (auction *Auction) Validate() *internal_error.InternalError {
	if len(auction.ProductName) <= 1 ||
		len(auction.Category) <= 2 ||
		len(auction.Description) <= 10 &&
			(auction.Condition != New && auction.Condition != Used && auction.Condition != Refurbished) {
		return internal_error.NewBadRequestError("invalid auction object")
	}
	return nil
}

func (auction *Auction) IsActive() bool {
	return auction.Status == Active
}

func (auction *Auction) IsCompleted() bool {
	return auction.Status == Completed
}

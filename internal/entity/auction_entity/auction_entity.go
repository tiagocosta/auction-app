package auction_entity

import "time"

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

func (auction *Auction) IsActive() bool {
	return auction.Status == Active
}

func (auction *Auction) IsCompleted() bool {
	return auction.Status == Completed
}

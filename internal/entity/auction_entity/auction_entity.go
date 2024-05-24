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

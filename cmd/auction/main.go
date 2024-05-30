package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tiagocosta/auction-app/configuration/database/mongodb"
	"github.com/tiagocosta/auction-app/internal/infra/api/web/controller/auction_controller"
	"github.com/tiagocosta/auction-app/internal/infra/api/web/controller/bid_controller"
	"github.com/tiagocosta/auction-app/internal/infra/api/web/controller/user_controller"
	"github.com/tiagocosta/auction-app/internal/infra/database/auction"
	"github.com/tiagocosta/auction-app/internal/infra/database/bid"
	"github.com/tiagocosta/auction-app/internal/infra/database/user"
	"github.com/tiagocosta/auction-app/internal/usecase/auction_usecase"
	"github.com/tiagocosta/auction-app/internal/usecase/bid_usecase"
	"github.com/tiagocosta/auction-app/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("error trying to load env variables")
		return
	}

	dbConn, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionController := initDependencies(dbConn)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController,
) {
	auctionRepositoty := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepositoty)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepositoty, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))
	return
}

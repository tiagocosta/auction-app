package mongodb

import (
	"context"
	"os"

	"github.com/tiagocosta/auction-app/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDB := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("error trying to connect to mongodb", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("error trying to ping mongodb", err)
		return nil, err
	}

	return client.Database(mongoDB), nil
}

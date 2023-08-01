package mongodb

import (
	"context"
	"os"

	"github.com/danilocordeirodev/go-basics/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
	MONGODB_USERNAME = "MONGODB_USERNAME"
	MONGODB_PASSWORD = "MONGODB_PASSWORD"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error){
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_DATABASE)

	credentials := options.Credential{
		AuthSource: "admin",
		Username: os.Getenv(MONGODB_USERNAME),
		Password: os.Getenv(MONGODB_PASSWORD),
	}
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri).SetAuth(credentials))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("conseguiu connection established")

	return client.Database(mongodb_database), nil
}
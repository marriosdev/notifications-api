package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConfig struct {
	URI      string
	Database string
}

func newMongoConfig() mongoConfig {
	return mongoConfig{
		URI:      os.Getenv("MONGO_DB_URI"),
		Database: "membrou",
	}
}

func NewMongoDB() (*mongo.Client, *mongo.Database, error) {
	cfg := newMongoConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	db := client.Database(cfg.Database)
	return client, db, nil
}

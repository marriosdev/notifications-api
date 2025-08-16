package notification

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
	mongoDb *mongo.Database
}

func NewRepository(mongo *mongo.Database) *Repository {
	return &Repository{
		mongoDb: mongo,
	}
}

func (r Repository) GetAll(ctx context.Context) ([]Notification, error) {
	cur, err := r.mongoDb.Collection("notifications").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var notifications []Notification
	if err := cur.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}

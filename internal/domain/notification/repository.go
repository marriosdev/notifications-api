package notification

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r Repository) GetById(ctx context.Context, id string) (Notification, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	cur, err := r.mongoDb.Collection("notifications").Find(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var notification Notification
	for cur.Next(ctx) {
		if err := cur.Decode(&notification); err != nil {
			log.Fatal(err)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return notification, nil
}

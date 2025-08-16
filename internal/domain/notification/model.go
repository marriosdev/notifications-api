package notification

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type        string             `bson:"type" json:"type"`
	Subject     string             `bson:"subject" json:"subject"`
	Message     string             `bson:"message" json:"message"`
	Status      string             `bson:"status" json:"status"`
	Destination string             `bson:"destination" json:"destination"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

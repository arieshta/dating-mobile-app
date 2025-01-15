package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Premium struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UserId         string             `bson:"user_id"`
	Duration       int                `bson:"premium"`
	PurchasedAt    time.Time          `bson:"purchased_at"`
	PurchaseNumber int                `bson:"purchase_number"`
}

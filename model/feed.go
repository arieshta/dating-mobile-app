package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feed struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID    string               `json:"user_id" bson:"user_id"`
	Feeds     []primitive.ObjectID `json:"feeds" bson:"feeds"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at"`
}

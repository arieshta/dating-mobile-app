package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feed struct {
	ID     primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID string               `json:"used_id" bson:"user_id"`
	Feeds  []primitive.ObjectID `json:"feeds" bson:"feeds"`
}
